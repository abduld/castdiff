// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This grammar is derived from the C grammar in the 'ansitize'
// program, which carried this notice:
//
// Copyright (c) 2006 Russ Cox,
// 	Massachusetts Institute of Technology
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the
// Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute,
// sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall
// be included in all copies or substantial portions of the
// Software.
//
// The software is provided "as is", without warranty of any
// kind, *SymbolLiteraless or implied, including but not limited to the
// warranties of merchantability, fitness for a particular
// purpose and noninfringement.  In no event shall the authors
// or copyright holders be liable for any claim, damages or
// other liability, whether in an action of contract, tort or
// otherwise, arising from, out of or in connection with the
// software or the use or other dealings in the software.

%{
package cc

import (
	//"runtime/debug"
	_ "fmt"

	. "github.com/abduld/castdiff/cc/ast"
)

type typeClass struct {
	c Storage
	q TypeQual
	t *Type
}

type idecor struct {
	d func(*Type) (*Type, *SymbolLiteral)
	i Expr
}

var id int = 0

func nextId() int {
	id += 1
	return id
}

%}

%union {
	abdecor func(*Type) *Type
	fundecl *FuncStmt
	decl Stmt
	decls []DeclStmt
	decor func(*Type) (*Type, *SymbolLiteral)
	decors []func(*Type) (*Type, *SymbolLiteral)
	expr Expr
	exprs []Expr
	idec idecor
	idecs []idecor
	init Expr
	inits []Expr
	label *Label
	labels []*Label
	span Span
	prefix Expr
	prefixes []Expr
	stmt Stmt
	stmts []Stmt
	str string
	strs []string
	tc typeClass
	tk TypeType
	typ *Type
	symlit *SymbolLiteral
	boollit *BooleanLiteral
	reallit *RealLiteral
	intlit *IntegerLiteral
	charlit *CharLiteral
	strlit *StringLiteral
	syntax Syntax
	syntaxs []Syntax
	langkey *LanguageKeyword
}

%token	<str>	tokARGBEGIN
%token	<str>	tokARGEND
%token	<str>	tokAUTOLIB
%token	<str>	tokSET
%token	<str>	tokUSED

%token	<str>	tokAuto
%token	<str>	tokBreak
%token	<str>	tokCase
%token	<str>	tokChar
%token	<str>	tokConst
%token	<str>	tokContinue
%token	<str>	tokDefault
%token	<str>	tokDo
%token	<str>	tokDotDotDot
%token	<str>	tokDouble
%token	<str>	tokEnum
%token	<str>	tokError
%token	<str>	tokExtern
%token	<str>	tokFor
%token	<str>	tokGoto
%token	<str>	tokIf
%token	<str>	tokInline
%token	<str>	tokName
%token	<intlit>	tokInteger
%token	<reallit>	tokReal
%token	<intlit>	tokInt
%token	<reallit>	tokFloat
%token	<charlit>	tokLitChar
%token	<intlit>	tokLong
%token	<strlit>	tokString
%token	<str>	tokOffsetof
%token	<str>	tokRegister
%token	<str>	tokReturn
%token	<str>	tokShort
%token	<str>	tokSigned
%token	<str>	tokStatic
%token	<str>	tokStruct
%token	<str>	tokSwitch
%token	<str>	tokTypeName
%token	<str>	tokTypedef
%token	<str>	tokUnion
%token	<str>	tokUnsigned
%token	<str>	tokVaArg
%token	<str>	tokVoid
%token	<str>	tokVolatile
%token	<str>	tokWhile

%token  <str>   tokLCuBrk
%token  <str>   tokRCuBrk

%token	<str>	tokDevice
%token	<str>	tokHost
%token	<str>	tokGlobal
%token	<str>	tokShared
%token	<str>	tokRestrict

%type	<abdecor>	abdecor abdec1
%type	<stmt>	fndef
%type	<stmt> edecl fnarg
%type	<stmts>	decl decl_list_opt
%type	<stmts>	fnarg_list fnarg_list_opt
%type	<stmts>	prog xdecl topdecl
%type	<stmts>	sudecl sudecl_list
%type	<stmts> edecl_list
%type	<decor>	decor sudecor
%type	<decors>	sudecor_list sudecor_list_opt
%type	<expr>	expr expr_opt cexpr cexpr_opt eqexpr eqexpr_opt
%type	<exprs>	expr_list expr_list_opt
%type	<idec>	idecor
%type	<idecs>	idecor_list idecor_list_opt
%type	<init>	init binit
%type	<inits>	braced_init_list binit_list
%type	<label>	label
%type	<labels> label_list_opt
%type	<prefix>	initprefix
%type	<prefixes>	initprefix_list
%type	<stmt>	stmt block lstmt
%type	<stmts>	block1
%type <symlit> tag
%type	<syntax>	cname qname tname cqname cqtname tag_opt
%type	<syntaxs>	cqname_list cqname_list_opt 
%type	<syntaxs>	cqtname_list cqtname_list_opt
%type	<syntaxs>	qname_list qname_list_opt
%type	<exprs>	string_list
%type	<tc>	typeclass
%type	<tk>	structunion
%type	<typ>	abtype type typespec

// fake operators to resolve if/else ambiguity
%left	tokShift
%left	tokElse
%left	tokTypeName
%left	'{'
%left	tokName

// real operators - usual c precedence
%left	','
%right	'=' tokAddEq tokSubEq tokMulEq tokDivEq tokModEq tokLshEq tokRshEq tokAndEq tokXorEq tokOrEq
%right	'?' ':'
%left	tokOrOr
%left	tokAndAnd
%left	'|'
%left	'^'
%left	'&'
%left	tokEqEq
%left	tokNotEq
%left	'<' '>' tokLtEq tokGtEq
%left	tokLsh tokRsh
%left	'+' '-'
%left	'*' '/' '%'
%right	tokCast
%left	'!' '~' tokSizeof tokUnary
%right	'.' '[' ']' ')' tokDec tokInc tokArrow
%right tokLCuBrk tokRCuBrk
%left	tokString

%left '('
%token startProg	startExpr tokEOF

%%


top:
	startProg prog tokEOF
	{
		yylex.(*lexer).prog = &Prog{Decls: $2, Id: nextId()}
		return 0
	}
|	startExpr cexpr tokEOF
	{
		return 0
	}

prog:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	prog xdecl
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = append($1, $2...)
	}
|	prog tokAUTOLIB '(' tokName ')'
	{
	}

cexpr:
	expr_list
	{
		$<span>$ = $<span>1
		if len($1) == 1 {
			$$ = $1[0]
			break
		}
		$$ = &CallExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$}, Callee: &SymbolLiteral{Value: ",",}, Args: $1,}
	}

expr:
	tokName
	{
		$<span>$ = $<span>1
		$$ = &SymbolLiteral{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	tokInteger
	{
		$<span>$ = $<span>1
		$$ = &IntegerLiteral{
			Value: $1.Value,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	tokReal
	{
		$<span>$ = $<span>1
		$$ = &RealLiteral{
			Value: $1.Value,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	tokLitChar
	{
		$<span>$ = $<span>1
		$$ = &CharLiteral{
			Value: $1.Value,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	string_list
	{
		$<span>$ = $<span>1
		$$ = &TupleExpr{
			Args: $1,
			Brackets: None,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	expr '+' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Add,
			Left: $1,
			Right: $3,
		}
	}
|	expr '-' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Sub,
			Left: $1,
			Right: $3,
		}
	}
|	expr '*' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Mul,
			Left: $1,
			Right: $3,
		}
	}
|	expr '/' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Div,
			Left: $1,
			Right: $3,
		}
	}
|	expr '%' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Mod,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokLsh expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Lsh,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokRsh expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Rsh,
			Left: $1,
			Right: $3,
		}
	}
|	expr '<' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Lt,
			Left: $1,
			Right: $3,
		}
	}
|	expr '>' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Gt,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokLtEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: LtEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokGtEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: GtEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokEqEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: EqEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokNotEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: NotEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr '&' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: And,
			Left: $1,
			Right: $3,
		}
	}
|	expr '^' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Xor,
			Left: $1,
			Right: $3,
		}
	}
|	expr '|' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Or,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokAndAnd expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: AndAnd,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokOrOr expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BinaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: OrOr,
			Left: $1,
			Right: $3,
		}
	}
|	expr '?' cexpr ':' expr
	{
		$<span>$ = span($<span>1, $<span>5)

		$$ = &CondExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Cond: $1,
			Then: $3,
			Else: $5,
		}
	}
|	expr '=' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Eq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokAddEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: AddEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokSubEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: SubEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokMulEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: MulEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokDivEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: DivEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokModEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: ModEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokLshEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: LshEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokRshEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: RshEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokAndEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: AndEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokXorEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: XorEq,
			Left: $1,
			Right: $3,
		}
	}
|	expr tokOrEq expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &AssignExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: OrEq,
			Left: $1,
			Right: $3,
		}
	}
|	'*' expr	%prec tokUnary
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &UnaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Indir,
			Arg: $2,
		}
	}
|	'&' expr	%prec tokUnary
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &UnaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Addr,
			Arg: $2,
		}
	}
|	'+' expr	%prec tokUnary
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &UnaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Plus,
			Arg: $2,
		}
	}
|	'-' expr	%prec tokUnary
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &UnaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Minus,
			Arg: $2,
		}
	}
|	'!' expr
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &UnaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Not,
			Arg: $2,
		}
	}
|	'~' expr
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &UnaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: Twid,
			Arg: $2,
		}
	}
|	tokInc expr
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &UnaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: PreInc,
			Arg: $2,
		}
	}
|	tokDec expr
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &UnaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Op: PreDec,
			Arg: $2,
		}
	}
|	tokSizeof expr	%prec tokSizeof
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &CallExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Callee: &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Id: nextId(),
				Value: "sizeof",
			},
			Args: []Expr{$2},
		}
	}
|	tokSizeof '(' abtype ')'	%prec tokSizeof
	{
		$<span>$ = span($<span>1, $<span>4)
		$$ = &CallExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Callee: &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Id: nextId(),
				Value: "sizeof", // sizeof type
			},
			Args: []Expr{$3},
		}
	}
|	tokOffsetof '(' abtype ',' expr ')'	%prec tokSizeof
	{
		$<span>$ = span($<span>1, $<span>6)
		$$ = &CallExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Callee: &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Id: nextId(),
				Value: "offsetof",
			},
			Args: []Expr{$3, $5},
		}
	}
|	'(' abtype ')' expr	%prec tokCast
	{
		$<span>$ = span($<span>1, $<span>4)
		$$ = &CastExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Type: $2,
			Expr: $4,
		}
	}
|	'(' abtype ')' braced_init_list	%prec tokCast
	{
		$<span>$ = span($<span>1, $<span>4)
		$$ = &CastExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Type: $2,
			Expr: &TupleExpr{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Brackets: Brace,
				Args: $4,
				Id: nextId(),
			},
		}
	}
|	'(' cexpr ')'
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &ParenExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Expr: $2,
			Id: nextId(),
		}
	}
|	expr tokLCuBrk expr_list_opt tokRCuBrk '(' expr_list_opt ')'
	{

		$<span>$ = span($<span>1, $<span>7)
		$$ = &CUDALaunchExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Callee: $1,
			LaunchParams: $3,
			Args: $6,
			Id: nextId(),
		}
	}
|	expr '(' expr_list_opt ')'
	{
		$<span>$ = span($<span>1, $<span>4)
		$$ = &CallExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Callee: $1,
			Args: $3,
			Id: nextId(),
		}
	}
|	expr '[' cexpr ']'
	{
		$<span>$ = span($<span>1, $<span>4)
		$$ = &CallExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Callee: &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Id: nextId(),
				Value: "array_index",
			},
			Args: []Expr{$1, $3},
			Id: nextId(),
		}
	}
|	expr tokInc
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &UnaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Op: PostInc,
			Arg: $1,
			Id: nextId(),
		}
	}
|	expr tokDec
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &UnaryExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Op: PostDec,
			Arg: $1,
			Id: nextId(),
		}
	}
|	tokVaArg '(' expr ',' abtype ')'
	{
		$<span>$ = span($<span>1, $<span>6)
		$$ = &VaArgExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Type: $5,
			Arg: $3,
			Id: nextId(),
		}
	}

block1:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	block1 decl
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = $1
		for _, d := range $2 {
			$$ = append($$, d)
		}
	}
|	block1 lstmt
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = append($1, $2)
	}

block:
	'{'
	{
		yylex.(*lexer).pushScope()
	}
	block1 '}'
	{
		$<span>$ = span($<span>1, $<span>4)
		yylex.(*lexer).popScope()
		$$ = &BlockStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Stmts: $3}
	}

label:
	tokCase expr ':'
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &Label{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Name: $2, IsCase: true }
	}
|	tokDefault ':'
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &Label{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Name: &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Id: nextId(),
				Value: "default",
			},
		}
	}
|	tokName ':'
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &Label{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Name: &SymbolLiteral{
				Id: nextId(),
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Value: $1,
			},
		}
	}

lstmt:
	label_list_opt stmt
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = $2
		$$ = &LabeledStmt{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Labels: $1,
			Expr: $2,
		}
	}

stmt:
	';'
	{
		$<span>$ = $<span>1
		$$ = &EmptyStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId()}
	}
|	tokUSED '(' cexpr ')' ';'
	{
		$<span>$ = $<span>1
		$$ = &EmptyStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId()}
	}
|	tokSET '(' cexpr ')' ';'
	{
		$<span>$ = $<span>1
		$$ = &EmptyStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId()}
	}
|	block
	{
		$<span>$ = $<span>1
		$$ = $1
	}
|	cexpr ';'
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &ExprStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Expr: $1}
	}
|	tokARGBEGIN block1 tokARGEND
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &BlockStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Stmts: $2}
	}
|	tokBreak ';'
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &BreakStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId()}
	}
|	tokContinue ';'
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &ContinueStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId()}
	}
|	tokDo lstmt tokWhile '(' cexpr ')' ';'
	{
		$<span>$ = span($<span>1, $<span>7)
		$$ = &DoStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Body: $2, Cond: $5}
	}
|	tokFor '(' cexpr_opt ';' cexpr_opt ';' cexpr_opt ')' lstmt
	{
		$<span>$ = span($<span>1, $<span>9)
		$$ = &ForStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Init: $3,
			Cond: $5,
			Update: $7,
			Body: $9,
		}
	}
|	tokGoto tag ';'
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &GotoStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Target: $2}
	}
|	tokIf '(' cexpr ')' lstmt	%prec tokShift
	{
		$<span>$ = span($<span>1, $<span>5)
		$$ = &IfStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Cond: $3, Then: $5}
	}
|	tokIf '(' cexpr ')' lstmt tokElse lstmt
	{
		$<span>$ = span($<span>1, $<span>7)
		$$ = &IfStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Cond: $3, Then: $5, Else: $7}
	}
|	tokReturn cexpr_opt ';'
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &ReturnStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Value: $2}
	}
|	tokSwitch '(' cexpr ')' lstmt
	{
		$<span>$ = span($<span>1, $<span>5)
		$$ = &SwitchStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Cond: $3, Body: $5}
	}
|	tokWhile '(' cexpr ')' lstmt
	{
		$<span>$ = span($<span>1, $<span>5)
		$$ = &WhileStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Id: nextId(), Cond: $3, Body: $5}
	}

// Abstract declarator - abdec1 includes the slot where the name would go
abdecor:
	{
		$<span>$ = Span{}
		$$ = func(t *Type) *Type { return t}
	}
|	'*' qname_list_opt abdecor
	{
		$<span>$ = span($<span>1, $<span>3)
		_, q, _ := splitTypeWords($2)
		abdecor := $3
		$$ = func(t *Type) *Type {
			return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Type: Ptr, Base: t, Qual: q, Id: nextId()})
		}
	}
|	abdec1
	{
		$<span>$ = $<span>1
		$$ = $1
	}

abdec1:
	abdec1 '(' fnarg_list_opt ')'
	{
		$<span>$ = span($<span>1, $<span>4)
		abdecor := $1
		decls := $3
		span := $<span>$
		for _, decl := range decls {
			switch decl := decl.(type) {
				default:
					break ;
				case *DeclStmt:
					t := decl.Type
					if t != nil {
						if t.Type == TypedefType && t.Base != nil {
							t = t.Base
						}
						if t.Type == Array {
							if t.Size == nil {
								t = t.Base
							}
							decl.Type = &Type{Type: Ptr, Base: t, Id: nextId()}
						}
					}
			}
		}
		$$ = func(t *Type) *Type {
			return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Type: Func, Base: t, Stmts: decls, Id: nextId()})
		}
	}
|	abdecor '[' expr_opt ']'
	{
		$<span>$ = span($<span>1, $<span>4)
		abdecor := $1
		span := $<span>$
		expr := $3
		$$ = func(t *Type) *Type {
			return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Type: Array, Base: t, Size: expr, Id: nextId()})
		}

	}
|	'(' abdecor ')'
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = $2
	}

// Concrete declarator
decor:
	tag
	{
		$<span>$ = $<span>1
		name := $1
		$$ = func(t *Type) (*Type, *SymbolLiteral) { return t, name }
	}
|	'*' qname_list_opt decor
	{
		$<span>$ = span($<span>1, $<span>3)
		_, q, _ := splitTypeWords($2)
		decor := $3
		span := $<span>$
		$$ = func(t *Type) (*Type, *SymbolLiteral) {
			return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Type: Ptr, Base: t, Qual: q, Id: nextId()})
		}
	}
|	'(' decor ')'
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = $2
	}
|	decor '(' fnarg_list_opt ')'
	{
		$<span>$ = span($<span>1, $<span>4)
		decor := $1
		decls := $3
		span := $<span>$
		$$ = func(t *Type) (*Type, *SymbolLiteral) {
			return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Type: Func, Base: t, Stmts: decls, Id: nextId()})
		}
	}
|	decor '[' expr_opt ']'
	{
		$<span>$ = span($<span>1, $<span>4)
		decor := $1
		span := $<span>$
		expr := $3
		$$ = func(t *Type) (*Type, *SymbolLiteral) {
			return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Type: Array, Base: t, Size: expr, Id: nextId()})
		}
	}

// Function argument
fnarg:
	tokName
	{
		$<span>$ = $<span>1
		$$ = &DeclStmt{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Name: &SymbolLiteral{
				Value: $1,
				Id: nextId(),
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
			},
			Id: nextId(),
		}
	}
|	type abdecor
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &DeclStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Type: $2($1), Id: nextId()}
	}
|	type decor
	{
		$<span>$ = span($<span>1, $<span>2)
		typ, name := $2($1)
		$$ = &DeclStmt{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Name: name, Type: typ, Id: nextId()}
	}
|	tokDotDotDot
	{
		$<span>$ = $<span>1
		$$ = &DeclStmt{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Name: &SymbolLiteral{
				Value: "...",
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Id: nextId(),
			},
			Id: nextId(),
		}
	}

// Initialized declarator
idecor:
	decor
	{
		$<span>$ = $<span>1
		$$ = idecor{$1, nil}
	}
|	decor '=' init
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = idecor{$1, $3}
	}

// Class words
cname:
	tokAuto
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	tokStatic
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	tokExtern
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	tokTypedef
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	tokRegister
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	tokInline
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}

// Qualifier words
qname:
	tokConst
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
|	tokVolatile
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
	}
| tokDevice
  {
		$<span>$ = $<span>1
		$$ = &SymbolLiteral{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
  }
| tokHost
  {
		$<span>$ = $<span>1
		$$ = &SymbolLiteral{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
  }
| tokGlobal
  {
		$<span>$ = $<span>1
		$$ = &SymbolLiteral{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
  }
| tokShared
  {
		$<span>$ = $<span>1
		$$ = &SymbolLiteral{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
  }
| tokRestrict
  {
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			Value: $1,
			Id: nextId(),
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
		}
  }

// Type words
tname:
	tokChar
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: $1,
			Id: nextId(),
		}
	}
|	tokShort
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: $1,
			Id: nextId(),
		}
	}
|	tokInt
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: "int",
			Id: nextId(),
		}
	}
|	tokLong
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: "long",
			Id: nextId(),
		}
	}
|	tokSigned
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: "signed",
			Id: nextId(),
		}
	}
|	tokUnsigned
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: "unsigned",
			Id: nextId(),
		}
	}
|	tokFloat
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: "float",
			Id: nextId(),
		}
	}
|	tokDouble
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: "real",
			Id: nextId(),
		}
	}
|	tokVoid
	{
		$<span>$ = $<span>1
		$$ = &LanguageKeyword{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: "void",
			Id: nextId(),
		}
	}

cqname:
	cname
	{
		$<span>$ = $<span>1
		$$ = $1
	}
|	qname
	{
		$<span>$ = $<span>1
		$$ = $1
	}

cqtname:
	cqname
	{
		$<span>$ = $<span>1
		$$ = $1
	}
|	tname
	{
		$<span>$ = $<span>1
		$$ = $1
	}

// Type specifier but not a tname
typespec:
	tokTypeName
	{
		$<span>$ = $<span>1
		$$ = $<typ>1
		if $$ == nil {
			$$ = &Type{
				Type: TypedefType,
				Name: &SymbolLiteral{
					Value:$<str>1,
					Id: nextId(),
					SyntaxInfo: SyntaxInfo{Span: $<span>$},
				},
				Id: nextId(),
			}
		}
	}

// Types annotated with class info.
//	typeclass:
//		cqname* typespec cqname*
//	|	cqname* tname cqtname*
//	|	cqname+
// except LALR(1) can't handle that.
typeclass:
	cqname_list %prec tokShift
	{
		$<span>$ = $<span>1
		$$.c, $$.q, $$.t = splitTypeWords(
			append($1, &LanguageKeyword{
				Value: "int",
				Id: nextId(),
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
			}))
	}
|	cqname_list typespec cqname_list_opt
	{
		$<span>$ = span($<span>1, $<span>3)
		$$.c, $$.q, _ = splitTypeWords(append($1, $3...))
		$$.t = $2
	}
|	cqname_list tname cqtname_list_opt
	{
		$<span>$ = span($<span>1, $<span>3)
		$1 = append($1, $2)
		$1 = append($1, $3...)
		$$.c, $$.q, $$.t = splitTypeWords($1)
	}
|	typespec cqname_list_opt
	{
		$<span>$ = span($<span>1, $<span>2)
		$$.c, $$.q, _ = splitTypeWords($2)
		$$.t = $1
	}
|	tname cqtname_list_opt
	{
		$<span>$ = span($<span>1, $<span>2)
		var ts []Syntax
		ts = append(ts, $1)
		ts = append(ts, $2...)
		//PrintStack()
		$$.c, $$.q, $$.t = splitTypeWords(ts)
	}

// Types without class info (check for class in higher level)
type:
	typeclass
	{
		$<span>$ = $<span>1
		if $1.c != 0 {
			yylex.(*lexer).Errorf("%v not allowed here", $1.c)
		}
		if $1.q != 0 && $1.q != Const && $1.q != Volatile {
			yylex.(*lexer).Errorf("%v ignored here (TODO)?", $1.q)
		}
		$$ = $1.t
	}

abtype:
	type abdecor
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = $2($1)
	}

// Declaration (finally)
decl:
	typeclass idecor_list_opt ';'
	{
		lx := yylex.(*lexer)
		$<span>$ = span($<span>1, $<span>3)
		// TODO: use $1.q
		$$ = nil
		for _, idec := range $2 {
			typ, name := idec.d($1.t)
			d := &DeclStmt{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Name: name,
				Type: typ,
				Storage: $1.c,
				Init: idec.i,
				Id: nextId(),
			}
			lx.pushDecl(d);
			$$ = append($$, d);
		}
		if $2 == nil {
			d := &DeclStmt{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Name: &SymbolLiteral{},
				Type: $1.t,
				Storage: $1.c,
				Id: nextId(),
			}
			lx.pushDecl(d);
			$$ = append($$, d)
		}
	}

topdecl:
	typeclass idecor_list_opt ';'
	{
		lx := yylex.(*lexer)
		$<span>$ = span($<span>1, $<span>3)
		// TODO: use $1.q
		$$ = nil
		for _, idec := range $2 {
			typ, name := idec.d($1.t)
			d := lx.lookupDecl(name)
			if d == nil {
				d = &DeclStmt{
					SyntaxInfo: SyntaxInfo{Span: $<span>$},
					Name: name,
					Type: typ,
					Storage: $1.c,
					Init: idec.i,
					Id: nextId(),
				}
				lx.pushDecl(d)
			} else {
				//d.Span = $<span>$
				if idec.i != nil {
					switch d := d.(type) {
					case *DeclStmt:
						d.Init = idec.i
					}
				}
			}
			$$ = append($$, d);
		}
		if $2 == nil {
			d := &DeclStmt{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Name: &SymbolLiteral{},
				Type: $1.t,
				Storage: $1.c,
				Id: nextId(),
			}
			lx.pushDecl(d);
			$$ = append($$, d)
		}
	}

xdecl:
	topdecl
	{
		$<span>$ = $<span>1
		$$ = $1
	}
|	fndef
	{
		$<span>$ = $<span>1
		$$ = []Stmt{$1}
	}
|	tokExtern tokString '{' prog '}'
	{
		$$ = $4
	}

fndef:
	typeclass decor decl_list_opt
	{
		lx := yylex.(*lexer)
		typ, name := $2($1.t)
		if typ.Type != Func {
			yylex.(*lexer).Errorf("invalid function definition")
			return 0
		}
		d := lx.lookupDecl(name)
		if d == nil {
			d = &DeclStmt{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Name: name,
				Type: typ,
				Storage: $1.c,
				Id: nextId(),
			}
			lx.pushDecl(d);
		} else {
			switch d := d.(type) {
				case *DeclStmt:
					d.Type = typ
			}
		}
		$<decl>$ = d
		lx.pushScope()
		for _, decl := range typ.Stmts {
			lx.pushDecl(decl);
		}
	}
	block
	{
		yylex.(*lexer).popScope();
		$<span>$ = span($<span>1, $<span>5)
		$$ = $<decl>4
		if $3 != nil {
			yylex.(*lexer).Errorf("cannot use pre-prototype definitions")
		}
		switch decl := $$.(type) {
			case *DeclStmt:
				$$ = &FuncStmt{
					SyntaxInfo: SyntaxInfo{Span: $<span>$},
					Name: decl.Name,
					ReturnType: decl.Type,
					IsDecl: false,
					Storage: decl.Storage,
					Body: $5,
				}
			case *FuncStmt:
				$$ = &FuncStmt{
					SyntaxInfo: SyntaxInfo{Span: $<span>$},
					Name: decl.Name,
					ReturnType: decl.ReturnType,
					Args: decl.Args,
					IsDecl: false,
					Storage: decl.Storage,
					Body: $5,
				}
		}
	}

tag:
	tokName
	{
		$<span>$ = $<span>1
		$$ = &SymbolLiteral{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: $1,
			Id: nextId(),
		}
	}
|	tokTypeName
	{
		$<span>$ = $<span>1
		$$ = &SymbolLiteral{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Value: $1,
			Id: nextId(),
		}
	}

// struct/union
structunion:
	tokStruct
	{
		$<span>$ = $<span>1
		$$ = Struct
	}
|	tokUnion
	{
		$<span>$ = $<span>1
		$$ = Union
	}

sudecor:
	decor
	{
		$<span>$ = $<span>1
		$$ = $1
	}
|	tag_opt ':' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		name := $1
		expr := $3
		$$ = func(t *Type) (*Type, *SymbolLiteral) {
			t.Size = expr
			switch name := name.(type) {
				case *SymbolLiteral:
					return t, name
				default:
					return t, &SymbolLiteral{
						SyntaxInfo: SyntaxInfo{Span: $<span>$},
						Value: name.String(),
						Id: nextId(),
					}
			}
		}
	}

sudecl:
	type sudecor_list_opt ';'
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = nil
		for _, decor := range $2 {
			typ, name := decor($1)
			$$ = append($$, &DeclStmt{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Name: name,
				Type: typ,
				Id: nextId(),
			})
		}
		if $2 == nil {
			$$ = append($$, &DeclStmt{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Type: $1,
				Id: nextId(),
			})
		}
	}

typespec:
	structunion tag
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = yylex.(*lexer).pushType(&Type{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Type: $1,
			Tag: $2,
			Id: nextId(),
		})
	}
|	structunion tag_opt '{' sudecl_list '}'
	{
		$<span>$ = span($<span>1, $<span>5)
		$$ = yylex.(*lexer).pushType(&Type{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Type: $1,
			Tag: $2,
			Stmts: $4,
			Id: nextId(),
		})
	}

initprefix:
	'.' tag
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = &DotExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Struct: &EmptyLiteral{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Id: nextId(),
			},
			Member: $2,
		}
	}

expr:
	expr tokArrow tag
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &ArrowExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Ptr: $1,
			Member: $3,
		}
	}
|	expr '.' tag
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &DotExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Id: nextId(),
			Struct: $1,
			Member: $3,
		}
	}

// enum
typespec:
	tokEnum tag
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Type: Enum, Tag: $2, Id: nextId()})
	}
|	tokEnum tag_opt '{' edecl_list comma_opt '}'
	{
		$<span>$ = span($<span>1, $<span>6)
		$$ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Type: Enum, Tag: $2, Stmts: $4, Id: nextId()})
	}

edecl:
	tokName eqexpr_opt
	{
		$<span>$ = span($<span>1, $<span>2)
		var x *Init
		if $2 != nil {
			x = &Init{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Expr: $2, Id: nextId()}
		}
		$$ = &DeclStmt{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Name: &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Value: $1,
				Id: nextId(),
			},
			Init: x,
			Id: nextId(),
		}
		yylex.(*lexer).pushDecl($$);
	}

eqexpr:
	'=' expr
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = $2
	}

// initializers
init:
	expr
	{
		$<span>$ = $<span>1
		$$ = &Init{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Expr: $1, Id: nextId()}
	}
|	braced_init_list
	{
		$<span>$ = $<span>1
		$$ = &Init{SyntaxInfo: SyntaxInfo{Span: $<span>$}, Braced: $1, Id: nextId()}
	}

braced_init_list:
	'{' '}'
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = []Expr{
			&BracedExpr{SyntaxInfo: SyntaxInfo{Span: $<span>$},Id: nextId()},
		}
	}
|	'{' binit_list binit '}'
	{
		$<span>$ = span($<span>1, $<span>4)
		$$ = append($2, $3)
	}
|	'{' binit_list binit ',' '}'
	{
		$<span>$ = span($<span>1, $<span>5)
		$$ = append($2, $3)
	}

binit_list:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	binit_list binit ','
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = append($1, $2)
	}

binit:
	init
	{
		$<span>$ = $<span>1
		$$ = $1
	}
|	initprefix_list eq_opt init
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &TupleExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Brackets: None,
			Id: nextId(),
			Args: append($1, $3),
		}
	}

initprefix:
	'[' expr ']'
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = &CallExpr{
			SyntaxInfo: SyntaxInfo{Span: $<span>$},
			Callee: &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: $<span>$},
				Id: nextId(),
				Value: "array_index",
			},
			Args: []Expr{$2},
			Id: nextId(),
		}
	}

eq_opt:
	{
		$<span>$ = Span{}
	}
|	'='
	{
		$<span>$ = $<span>1
	}

comma_opt:
	{
		$<span>$ = Span{}
	}
|	','
	{
		$<span>$ = $<span>1
	}

// Special notations - should be created implicitly
// if we ever finish the yacc replacement.

initprefix_list:
	initprefix
	{
		$<span>$ = $<span>1
		$$ = []Expr{$1}
	}
|	initprefix_list initprefix
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = append($1, $2)
	}

tag_opt:
	{
		$<span>$ = Span{}
		$$ = &EmptyLiteral{}
	}
|	tag
	{
		$<span>$ = $<span>1
		$$ = $1
	}

cexpr_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	cexpr
	{
		$<span>$ = $<span>1
		$$ = $1
	}

expr_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	expr
	{
		$<span>$ = $<span>1
		$$ = $1
	}

expr_list:
	expr
	{
		$<span>$ = $<span>1
		$$ = []Expr{$1}
	}
|	expr_list ',' expr
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = append($1, $3)
	}

expr_list_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	expr_list
	{
		$<span>$ = $<span>1
		$$ = $1
	}

decl_list_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	decl_list_opt decl
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = append($1, $2...)
	}

label_list_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	label_list_opt label
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = append($1, $2)
	}

fnarg_list:
	fnarg
	{
		$<span>$ = $<span>1
		$$ = []Stmt{$1}
	}
|	fnarg_list ',' fnarg
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = append($1, $3)
	}

fnarg_list_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	fnarg_list
	{
		$<span>$ = $<span>1
		$$ = $1
	}

idecor_list:
	idecor
	{
		$<span>$ = $<span>1
		$$ = []idecor{$1}
	}
|	idecor_list ',' idecor
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = append($1, $3)
	}

idecor_list_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	idecor_list
	{
		$<span>$ = $<span>1
		$$ = $1
	}

qname_list:
	qname
	{
		$<span>$ = $<span>1
		$$ = []Syntax{$1}
	}
|	qname_list qname
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = append($1, $2)
	}

qname_list_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	qname_list
	{
		$<span>$ = $<span>1
		$$ = $1
	}

cqname_list:
	cqname
	{
		$<span>$ = $<span>1
		$$ = []Syntax{$1}
	}
|	cqname_list cqname
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = append($1, $2)
	}

cqname_list_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	cqname_list
	{
		$<span>$ = $<span>1
		$$ = $1
	}

cqtname_list:
	cqtname
	{
		$<span>$ = $<span>1
		$$ = []Syntax{$1}
	}
|	cqtname_list cqtname
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = append($1, $2)
	}

cqtname_list_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	cqtname_list
	{
		$<span>$ = $<span>1
		$$ = $1
	}

sudecor_list:
	sudecor
	{
		$<span>$ = $<span>1
		$$ = nil
		$$ = append($$, $1)
	}
|	sudecor_list ',' sudecor
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = append($1, $3)
	}

sudecor_list_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	sudecor_list
	{
		$<span>$ = $<span>1
		$$ = $1
	}

sudecl_list:
	sudecl
	{
		$<span>$ = $<span>1
		$$ = $1
	}
|	sudecl_list sudecl
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = append($1, $2...)
	}

eqexpr_opt:
	{
		$<span>$ = Span{}
		$$ = nil
	}
|	eqexpr
	{
		$<span>$ = $<span>1
		$$ = $1
	}

edecl_list:
	edecl
	{
		$<span>$ = $<span>1
		$$ = []Stmt{$1}
	}
|	edecl_list ',' edecl
	{
		$<span>$ = span($<span>1, $<span>3)
		$$ = append($1, $3)
	}

string_list:
	tokString
	{
		$<span>$ = $<span>1
		if $1 == nil {
			$$ = []Expr{
				&StringLiteral{
					Value: "",
					Id: nextId(),
					SyntaxInfo: SyntaxInfo{Span: $<span>$},
				},
			}
		} else {
			$$ = []Expr{
				&StringLiteral{
					Value: $1.Value,
					Id: nextId(),
					SyntaxInfo: SyntaxInfo{Span: $<span>$},
				},
			}
		}
	}
|	string_list tokString
	{
		$<span>$ = span($<span>1, $<span>2)
		$$ = append($1, &StringLiteral{
				SyntaxInfo: SyntaxInfo{Span: $<span>2},
				Value: $2.Value,
				Id: nextId(),
		})
	}


