// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go tool yacc cc.y

package cc

import (
	"fmt"
	. "github.com/abduld/castdiff/cc/ast"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Scope struct {
	Decl map[string]Stmt
	Tag  map[string]*Type
	Next *Scope
}

type lexer struct {
	// input
	start int
	byte  int
	lexInput
	pushed   []lexInput
	forcePos Position
	comments []Comment

	// comment assignment
	pre      []Syntax
	post     []Syntax
	enumSeen map[interface{}]bool

	// type checking state
	scope       *Scope
	includeSeen map[string]*Header

	// output
	errors []string
	prog   *Prog
	expr   *Expr
}

type Header struct {
	decls []Stmt
	types []*Type
}

func pushDeclStmt(lx *lexer, decl * DeclStmt) {
	sc := lx.scope

	if sc == nil {
		panic("no scope")
	}
	if decl.Name == nil {
		return
	}
	if sc.Decl == nil {
		sc.Decl = make(map[string]Stmt)
	}
	sc.Decl[decl.Name.Value] = decl
	if hdr := lx.declSave; hdr != nil && sc.Next == nil {
		hdr.decls = append(hdr.decls, decl)
	}
}

func pushFuncStmt(lx *lexer, decl * FuncStmt) {
	sc := lx.scope

	if sc == nil {
		panic("no scope")
	}
	if decl.Name == nil {
		return
	}
	if sc.Decl == nil {
		sc.Decl = make(map[string]Stmt)
	}
	sc.Decl[decl.Name.Value] = decl
	if hdr := lx.declSave; hdr != nil && sc.Next == nil {
		hdr.decls = append(hdr.decls, decl)
	}
}
func (lx *lexer) pushDecl(decl Stmt) {
	switch decl := decl.(type) {
	default:
		panic("undefined decl")
	case *DeclStmt:
		pushDeclStmt(lx, decl)
	case *FuncStmt:
		pushFuncStmt(lx, decl)
	}
}

func (lx *lexer) lookupDecl(sym *SymbolLiteral) Stmt {
	if sym == nil {
		return nil
	}
	name := sym.Value
	for sc := lx.scope; sc != nil; sc = sc.Next {
		decl := sc.Decl[name]
		if decl != nil {
			return decl
		}
	}
	return nil
}

func (lx *lexer) pushType(typ *Type) *Type {
	sc := lx.scope
	if sc == nil {
		panic("no scope")
	}

	if typ.Type == Enum && typ.Stmts != nil {
		for _, decl := range typ.Stmts {

			switch decl := decl.(type) {
			default:
				break;
			case *DeclStmt:
				lx.pushDecl(decl)
			case *FuncStmt:
				lx.pushDecl(decl)
			}
		}
	}

	if typ.Tag.String() == "" {
		return typ
	}

	old := lx.lookupTag(typ.Tag.String())
	if old == nil {
		if sc.Tag == nil {
			sc.Tag = make(map[string]*Type)
		}
		sc.Tag[typ.Tag.String()] = typ
		if hdr := lx.declSave; hdr != nil && sc.Next == nil {
			hdr.types = append(hdr.types, typ)
		}
		return typ
	}

	// merge typ into old
	if old.Kind != typ.Kind {
		lx.Errorf("conflicting tags: %s %s and %s %s", old.Kind, old.Tag, typ.Kind, typ.Tag)
		return typ
	}
	if typ.Stmts != nil {
		if old.Stmts != nil {
			lx.Errorf("multiple definitions for %s %s", old.Kind, old.Tag)
		}
		old.SyntaxInfo = typ.SyntaxInfo
		old.Stmts = typ.Stmts
	}
	return old
}

func (lx *lexer) lookupTag(name string) *Type {
	for sc := lx.scope; sc != nil; sc = sc.Next {
		typ := sc.Tag[name]
		if typ != nil {
			return typ
		}
	}
	return nil
}

func (lx *lexer) pushScope() {
	sc := &Scope{Next: lx.scope}
	lx.scope = sc
}

func (lx *lexer) popScope() {
	lx.scope = lx.scope.Next
}

func (lx *lexer) parse() {
	if lx.includeSeen == nil {
		lx.includeSeen = make(map[string]*Header)
	}
	if lx.wholeInput == "" {
		lx.wholeInput = lx.input
	}
	lx.scope = &Scope{}
	yyParse(lx)
}

type lexInput struct {
	wholeInput string
	input      string
	tok        string
	lastsym    string
	file       string
	lineno     int
	declSave   *Header
}

func (lx *lexer) pushInclude(includeLine string) {
	s := strings.TrimSpace(strings.TrimPrefix(includeLine, "#include"))
	if !strings.HasPrefix(s, "<") && !strings.HasPrefix(s, "\"") {
		lx.Errorf("malformed #include")
		return
	}
	sep := ">"
	if s[0] == '"' {
		sep = "\""
	}
	i := strings.Index(s[1:], sep)
	if i < 0 {
		lx.Errorf("malformed #include")
		return
	}
	i++

	file := s[1:i]
	origFile := file

	file, data, err := lx.findInclude(file, s[0] == '<')
	if err != nil {
		lx.Errorf("#include %s: %v", s[:i+1], err)
		return
	}

	if hdr := lx.includeSeen[file]; hdr != nil {
		for _, decl := range hdr.decls {
			// fmt.Printf("%s: replay %s\n", file, decl.Name)
			lx.pushDecl(decl)
		}
		for _, typ := range hdr.types {
			lx.pushType(typ)
		}
		return
	}

	hdr := lx.declSave
	if hdr == nil {
		hdr = new(Header)
		lx.includeSeen[file] = hdr
	} else {
		fmt.Printf("%s: warning nested %s\n", lx.span(), includeLine)
	}

	if extraMap[origFile] != "" {
		str := extraMap[origFile] + "\n"
		lx.pushed = append(lx.pushed, lx.lexInput)
		lx.lexInput = lexInput{
			input:      str,
			wholeInput: str,
			file:       "internal/" + origFile,
			lineno:     1,
			declSave:   hdr,
		}
	}

	if data == nil {
		return
	}

	lx.pushed = append(lx.pushed, lx.lexInput)
	str := string(append(data, '\n'))
	lx.lexInput = lexInput{
		input:      str,
		wholeInput: str,
		file:       file,
		lineno:     1,
		declSave:   hdr,
	}
}

var stdMap = map[string]string{
	"u.h":      hdr_u_h,
	"libc.h":   hdr_libc_h,
	"wb.h":     hdr_wb_h,
	"math.h": 	hdr_math_h,
	"stdarg.h": "",
	"signal.h": "",
}

var extraMap = map[string]string{}

var includes []string

func AddInclude(dir string) {
	includes = append(includes, dir)
}

func (lx *lexer) findInclude(name string, std bool) (string, []byte, error) {
	if std || name == "wb.h" || name == "wb" {
		if redir, ok := stdMap[name]; ok {
			if redir == "" {
				return "", nil, nil
			}
			return "internal/" + name, []byte(redir), nil
		}
		name = "/Users/abduld/" + name
	}
	if !filepath.IsAbs(name) {
		name1 := filepath.Join(filepath.Dir(lx.file), name)
		if _, err := os.Stat(name1); err != nil {
			for _, dir := range includes {
				name2 := filepath.Join(dir, name)
				if _, err := os.Stat(name2); err == nil {
					name1 = name2
					break
				}
			}
		}
		name = name1
	}
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return "", nil, err
	}
	return name, data, nil
}

func (lx *lexer) pop() bool {
	if len(lx.pushed) == 0 {
		return false
	}
	lx.lexInput = lx.pushed[len(lx.pushed)-1]
	lx.pushed = lx.pushed[:len(lx.pushed)-1]
	return true
}

func (lx *lexer) pos() Position {
	if lx.forcePos.Line != 0 {
		return lx.forcePos
	}
	return Position{
		File: lx.file,
		Line: lx.lineno,
		Byte: lx.byte,
	}
}
func (lx *lexer) span() Span {
	p := lx.pos()
	return Span{p, p}
}

func (lx *lexer) setSpan(s Span) {
	lx.forcePos = s.Start
}

func span(l1, l2 Span) Span {
	if l1.Start.Line == 0 {
		return l2
	}
	if l2.Start.Line == 0 {
		return l1
	}
	return Span{l1.Start, l2.End}
}

func (lx *lexer) skip(i int) {
	lx.lineno += strings.Count(lx.input[:i], "\n")
	lx.input = lx.input[i:]
	lx.byte += i
}

func (lx *lexer) token(i int) {
	lx.tok = lx.input[:i]
	lx.skip(i)
}

func (lx *lexer) sym(i int) {
	lx.token(i)
	lx.lastsym = lx.tok
}

func (lx *lexer) comment(i int) {
	var c Comment
	c.Span.Start = lx.pos()
	c.Text = lx.input[:i]
	j := len(lx.wholeInput) - len(lx.input)
	for j > 0 && (lx.wholeInput[j-1] == ' ' || lx.wholeInput[j-1] == '\t') {
		j--
	}
	if j > 0 && lx.wholeInput[j-1] != '\n' {
		c.Suffix = true
	}
	prefix := lx.wholeInput[j : len(lx.wholeInput)-len(lx.input)]
	lines := strings.Split(c.Text, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, prefix) {
			lines[i] = line[len(prefix):]
		}
	}
	c.Text = strings.Join(lines, "\n")

	lx.skip(i)
	c.Span.End = lx.pos()
	lx.comments = append(lx.comments, c)
}

func isalpha(c byte) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || c == '_' || c >= 0x80 || '0' <= c && c <= '9'
}

func isspace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\v' || c == '\f'
}

func (lx *lexer) setEnd(yy *yySymType) {
	yy.span.End = lx.pos()
}

func (lx *lexer) Lex(yy *yySymType) (yyt int) {
	//defer func() { println("tok", yy.str, yyt) }()
	if lx.start != 0 {
		tok := lx.start
		lx.start = 0
		return tok
	}
	*yy = yySymType{}
	defer lx.setEnd(yy)
Restart:
	yy.span.Start = lx.pos()
	in := lx.input
	if len(in) == 0 {
		if lx.pop() {
			goto Restart
		}
		return tokEOF
	}
	c := in[0]
	if isspace(c) {
		i := 1
		for i < len(in) && isspace(in[i]) {
			i++
		}
		lx.skip(i)
		goto Restart
	}

	i := 0
	switch c {
	case '#':
		i++
		for in[i] != '\n' {
			if in[i] == '\\' && in[i+1] == '\n' && i+2 < len(in) {
				i++
			}
			i++
		}
		str := in[:i]
		lx.skip(i)
		if strings.HasPrefix(str, "#include") {
			lx.pushInclude(str)
		}
		goto Restart

	case 'L':
		if in[1] != '\'' && in[1] != '"' {
			break // goes to alpha case after switch
		}
		i = 1
		fallthrough

	case '"', '\'':
		q := in[i]
		i++ // for the quote
		for ; in[i] != q; i++ {
			if in[i] == '\n' {
				what := "string"
				if q == '\'' {
					what = "character"
				}
				lx.Errorf("unterminated %s constant", what)
			}
			if in[i] == '\\' {
				i++
			}
		}
		i++ // for the quote
		lx.sym(i)
		yy.str = lx.tok
		if q == '"' {
			return tokString
		} else {
			return tokLitChar
		}

	case '.':
		if in[1] < '0' || '9' < in[1] {
			if in[1] == '.' && in[2] == '.' {
				lx.token(3)
				return tokDotDotDot
			}
			lx.token(1)
			return int(c)
		}
		fallthrough

	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		resTok := tokInteger
		for '0' <= in[i] && in[i] <= '9' || in[i] == '.' || 'A' <= in[i] && in[i] <= 'Z' || 'a' <= in[i] && in[i] <= 'z' {
			if in[i] == '.' {
				resTok = tokReal
			}
			i++
		}
		lx.sym(i)
		if resTok == tokInteger {
			ival, _ := strconv.Atoi(lx.tok)
			yy.intlit = &IntegerLiteral{Value: ival}
		} else {
			fval, _ := strconv.ParseFloat(lx.tok, 64)
			yy.reallit = &RealLiteral{Value: fval}
		}
		return resTok

	case '/':
		switch in[1] {
		case '*':
			i := 2
			for ; ; i++ {
				if i+2 <= len(in) && in[i] == '*' && in[i+1] == '/' {
					i += 2
					break
				}
				if i >= len(in) {
					lx.Errorf("unterminated /* comment")
					return tokError
				}
			}
			lx.comment(i)
			goto Restart

		case '/':
			for in[i] != '\n' {
				i++
			}
			lx.comment(i)
			if len(lx.input) >= 2 && lx.input[0] == '\n' && lx.input[1] == '\n' {
				lx.skip(1)
				lx.comment(0)
			}
			goto Restart
		}
		fallthrough

	case '~', '*', '(', ')', '[', ']', '{', '}', '?', ':', ';', ',', '%', '^', '!', '=', '<', '>', '+', '-', '&', '|':
		if c == '-' && in[1] == '>' {
			lx.token(2)
			return tokArrow
		}
		if in[1] == '=' && tokEq[c] != 0 {
			lx.token(2)
			return int(tokEq[c])
		}
		if in[1] == c && tokTok[c] != 0 {
			if in[2] == '=' && tokTokEq[c] != 0 {
				lx.token(3)
				return int(tokTokEq[c])
			}
			if tokCuBrk[in[2]] != 0 && in[2] == c {
				lx.token(3)
				return int(tokCuBrk[in[2]])
			}
			if tokCuBrk[c] != 0 && isspace(in[2]) {
				i := 2
				for i < len(in) && isspace(in[i]) {
					i++
				}
				if in[i] == c {
					lx.token(i + 1)
					//fmt.Println(string(in[i]), string(c), i, tokCuBrk[c])
					return int(tokCuBrk[c])
				}
			}
			lx.token(2)
			return int(tokTok[c])
		}
		lx.token(1)
		return int(c)
	}

	if isalpha(c) {
		for isalpha(in[i]) {
			i++
		}
		lx.sym(i)
		switch lx.tok {
		case "Adr":
			lx.tok = "Addr"
		case "union":
			lx.tok = "struct"
		}
		yy.str = lx.tok
		if t := tokId[lx.tok]; t != 0 {
			return int(t)
		}
		yy.decl = lx.lookupDecl(&SymbolLiteral{Value: lx.tok})
		if yy.decl != nil {
			switch decl := yy.decl.(type) {
			default:
				break;
			case *DeclStmt:
				if decl.Storage & Typedef != 0 {
					t := decl.Type
					for t.Type == TypedefType && t.Base != nil {
						t = t.Base
					}
					yy.typ = &Type{Type: TypedefType, Name: &SymbolLiteral{Value: yy.str}, Base: t, Stmts: []Stmt{yy.decl}}
					return tokTypeName
				}
			case *FuncStmt:
				if decl.Storage & Typedef != 0 {
					t := decl.ReturnType
					for t.Type == TypedefType && t.Base != nil {
						t = t.Base
					}
					yy.typ = &Type{Type: TypedefType, Name: &SymbolLiteral{Value: yy.str}, Base: t, Stmts: []Stmt{yy.decl}}
					return tokTypeName
				}
			}
		}
		if lx.tok == "EXTERN" {
			goto Restart
		}
		return tokName
	}

	lx.Errorf("unexpected input byte %#02x (%c)", c, c)
	return tokError
}

func (lx *lexer) Error(s string) {
	lx.Errorf("%s near %s", s, lx.lastsym)
}

func (lx *lexer) Errorf(format string, args ...interface{}) {
	lx.errors = append(lx.errors, fmt.Sprintf("%s: %s", lx.span(), fmt.Sprintf(format, args...)))
}

var tokEq = [256]int32{
	'*': tokMulEq,
	'/': tokDivEq,
	'+': tokAddEq,
	'-': tokSubEq,
	'%': tokModEq,
	'^': tokXorEq,
	'!': tokNotEq,
	'=': tokEqEq,
	'<': tokLtEq,
	'>': tokGtEq,
	'&': tokAndEq,
	'|': tokOrEq,
}

var tokTok = [256]int32{
	'<': tokLsh,
	'>': tokRsh,
	'=': tokEqEq,
	'+': tokInc,
	'-': tokDec,
	'&': tokAndAnd,
	'|': tokOrOr,
}

var tokTokEq = [256]int32{
	'<': tokLshEq,
	'>': tokRshEq,
}

var tokCuBrk = [256]int32{
	'<': tokLCuBrk,
	'>': tokRCuBrk,
}

var tokId = map[string]int32{
	"auto":     tokAuto,
	"break":    tokBreak,
	"case":     tokCase,
	"char":     tokChar,
	"const":    tokConst,
	"continue": tokContinue,
	"default":  tokDefault,
	"do":       tokDo,
	"double":   tokDouble,
	"else":     tokElse,
	"enum":     tokEnum,
	"extern":   tokExtern,
	"float":    tokFloat,
	"for":      tokFor,
	"goto":     tokGoto,
	"if":       tokIf,
	"inline":   tokInline,
	"int":      tokInt,
	"long":     tokLong,
	"offsetof": tokOffsetof,
	"register": tokRegister,
	"return":   tokReturn,
	"short":    tokShort,
	"signed":   tokSigned,
	"sizeof":   tokSizeof,
	"static":   tokStatic,
	"struct":   tokStruct,
	"switch":   tokSwitch,
	"typedef":  tokTypedef,
	"union":    tokUnion,
	"unsigned": tokUnsigned,
	"va_arg":   tokVaArg,
	"void":     tokVoid,
	"volatile": tokVolatile,
	"while":    tokWhile,

	"__device__": tokDevice,
	"__host__":   tokHost,
	"__global__": tokGlobal,
	"__shared__": tokShared,
	"restrict":   tokRestrict,

	"ARGBEGIN": tokARGBEGIN,
	"ARGEND":   tokARGEND,
	"AUTOLIB":  tokAUTOLIB,
	"USED":     tokUSED,
	"SET":      tokSET,
}

// Comment assignment.
// We build two lists of all subexpressions, preorder and postorder.
// The preorder list is ordered by start location, with outer expressions first.
// The postorder list is ordered by end location, with outer expressions last.
// We use the preorder list to assign each whole-line comment to the syntax
// immediately following it, and we use the postorder list to assign each
// end-of-line comment to the syntax immediately preceding it.

// enum walks the expression adding it and its subexpressions to the pre list.
// The order may not reflect the order in the input.
func (lx *lexer) enum(x Syntax) {
	switch x := x.(type) {
	default:
		panic(fmt.Errorf("order: unexpected type %T", x))
	case nil:
		return
	case *EmptyLiteral:
		//ok
	case *BooleanLiteral:
		//ok
	case *IntegerLiteral:
		//ok
	case *CharLiteral:
		//ok
	case *RealLiteral:
		//ok
	case *StringLiteral:
		//ok
	case *SymbolLiteral:
		//ok
	case *LanguageKeyword:
		//ok
	case *DotExpr:
		if x == nil {
			return
		}
		lx.enum(x.Struct)
		lx.enum(x.Member)
	case *LabeledStmt:
		if x == nil {
			return
		}
		for _, lbl := range x.Labels {
			lx.enum(lbl)
		}
		lx.enum(x.Expr)
	case *AssignExpr:
		if x == nil {
			return
		}
		lx.enum(x.Left)
		lx.enum(x.Right)
	case *CallExpr:
		if x == nil {
			return
		}
		lx.enum(x.Callee)
		for _, arg := range x.Args {
		lx.enum(arg)
	}
	case *TupleExpr:
		if x == nil {
			return
		}
		for _, arg := range x.Args {
			lx.enum(arg)
		}
	case *CastExpr:
		if x == nil {
			return
		}
		lx.enum(x.Expr)
	case *ReturnStmt:
		if x == nil {
			return
		}
		lx.enum(x.Value)
	case *CUDALaunchExpr:
		if x == nil {
			return
		}
		lx.enum(x.Callee)
		for _, arg := range x.LaunchParams {
			lx.enum(arg)
		}
		for _, arg := range x.Args {
			lx.enum(arg)
		}
	case *IfStmt:
		if x == nil {
			return
		}
		lx.enum(x.Cond)
		lx.enum(x.Then)
		lx.enum(x.Else)
	case *ExprStmt:
		if x == nil {
			return
		}
		lx.enum(x.Expr)
	case *ArrowExpr:
		if x == nil {
			return
		}
		lx.enum(x.Ptr)
		lx.enum(x.Member)
	case *BinaryExpr:
		if x == nil {
			return
		}
		lx.enum(x.Left)
		lx.enum(x.Right)
	case *UnaryExpr:
		if x == nil {
			return
		}
		lx.enum(x.Arg)
	case *FuncStmt:
		if x == nil {
			return
		}
		lx.enum(x.Name)
		for _, y := range x.Args {
			lx.enum(y)
		}
		lx.enum(x.Body)
	case *Init:
		if x == nil {
			return
		}
		lx.enum(x.Expr)
		for _, y := range x.Prefix {
			lx.enum(y)
		}
		for _, y := range x.Braced {
			lx.enum(y)
		}
	case *Prog:
		if x == nil {
			return
		}
		for _, y := range x.Decls {
			lx.enum(y)
		}
	case *BlockStmt:
		if x == nil {
			return
		}
		for _, y := range x.Stmts {
			lx.enum(y)
		}
	case *Label:
		lx.enum(x.Name)
		// ok
	case *DeclStmt:
		if x == nil {
			return
		}
		if lx.enumSeen[x] {
			return
		}
		lx.enumSeen[x] = true
		lx.enum(x.Type)
		lx.enum(x.Name)
		lx.enum(x.Init)
	case *Type:
		if x == nil {
			return
		}
		lx.enum(x.Base)
		lx.enum(x.Tag)
		lx.enum(x.Name)
		for _, y := range x.Stmts {
			lx.enum(y)
		}
		return // do not record type itself, just inner decls
	}
	lx.pre = append(lx.pre, x)
}

func (lx *lexer) order(prog *Prog) {
	lx.enumSeen = make(map[interface{}]bool)
	lx.enum(prog)
	sort.Sort(byStart(lx.pre))
	lx.post = make([]Syntax, len(lx.pre))
	copy(lx.post, lx.pre)
	sort.Sort(byEnd(lx.post))
}

type byStart []Syntax

func (x byStart) Len() int      { return len(x) }
func (x byStart) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x byStart) Less(i, j int) bool {
	if x[i] == nil && x[j] == nil {
		return true
	}
	if x[i] == nil {
		return false
	}
	if x[j] == nil {
		return true
	}
	pi := x[i].GetSpan()
	pj := x[j].GetSpan()
	// Order by start byte, leftmost first,
	// and break ties by choosing outer before inner.
	if pi.Start.Byte != pj.Start.Byte {
		return pi.Start.Byte < pj.Start.Byte
	}
	return pi.End.Byte > pj.End.Byte
}

type byEnd []Syntax

func (x byEnd) Len() int      { return len(x) }
func (x byEnd) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x byEnd) Less(i, j int) bool {
	pi := x[i].GetSpan()
	pj := x[j].GetSpan()
	// Order by end byte, leftmost first,
	// and break ties by choosing inner before outer.
	if pi.End.Byte != pj.End.Byte {
		return pi.End.Byte < pj.End.Byte
	}
	return pi.Start.Byte > pj.Start.Byte
}
