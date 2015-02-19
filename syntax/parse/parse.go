// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go tool yacc cc.y

package cc

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func Read(name string, r io.Reader) (*Prog, error) {
	return ReadMany([]string{name}, []io.Reader{r})
}

func ReadMany(names []string, readers []io.Reader) (*Prog, error) {
	lx := &lexer{}
	var prog *Prog
	for i, name := range names {
		if lx.includeSeen[name] != nil {
			continue
		}
		r := readers[i]
		data, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
		data = append(data, '\n')
		lx.start = startProg
		lx.lexInput = lexInput{
			input:  string(data),
			file:   name,
			lineno: 1,
		}
		lx.parse()
		if lx.errors != nil {
			return nil, fmt.Errorf("%v", lx.errors[0])
		}
		if prog == nil {
			prog = lx.prog
		} else {
			prog.Span.End = lx.prog.Span.End
			prog.Decls = append(prog.Decls, lx.prog.Decls...)
		}
		lx.prog = nil
		for sc := lx.scope; sc != nil; sc = sc.Next {
			for name, decl := range sc.Decl {
				if decl.Storage&Static != 0 || (decl.Storage&Typedef != 0 && strings.HasSuffix(decl.Span.Start.File, ".c")) {
					delete(sc.Decl, name)
				}
			}
			for name, typ := range sc.Tag {
				if strings.HasSuffix(typ.Span.Start.File, ".c") {
					delete(sc.Tag, name)
				}
			}
		}
	}
	lx.prog = prog
	lx.assignComments()
	if lx.errors != nil {
		return nil, fmt.Errorf("%v", strings.Join(lx.errors, "\n"))
	}

	removeDuplicates(lx.prog)

	return lx.prog, nil
}

func ParseProg(str string) (*Prog, error) {
	lx := &lexer{
		start: startProg,
		lexInput: lexInput{
			input:  str + "\n",
			file:   "<string>",
			lineno: 1,
		},
	}
	lx.parse()
	if lx.errors != nil {
		return nil, fmt.Errorf("parsing expression %#q: %v", str, lx.errors[0])
	}
	removeDuplicates(lx.prog)
	return lx.prog, nil
}

func ParseExpr(str string) (*Expr, error) {
	lx := &lexer{
		start: startExpr,
		lexInput: lexInput{
			input:  str + "\n",
			file:   "<string>",
			lineno: 1,
		},
	}
	lx.parse()
	if lx.errors != nil {
		return nil, fmt.Errorf("parsing expression %#q: %v", str, lx.errors[0])
	}
	return lx.expr, nil
}

type Prog struct {
	SyntaxInfo
	Id    int
	Decls []*Decl
}

func (x *Prog) GetId() int {
	return x.Id
}

func (x *Prog) String() string {
	var p Printer
	p.hideComments = true
	p.printProg(x)
	return p.String()
}

func (x *Prog) GetChildren() []Syntax {
	lst := []Syntax{}
	for _, elem := range x.Decls {
		lst = append(lst, elem)
	}
	return lst
}

// removeDuplicates drops the duplicated declarations
// caused by forward decls from prog.
// It keeps the _last_ of each given declaration,
// assuming that's the complete one.
// This heuristic tends to preserve something like
// source order.
// It would be defeated by someone writing a "forward"
// declaration following the real definition.
func removeDuplicates(prog *Prog) {
	count := map[*Decl]int{}
	for _, d := range prog.Decls {
		count[d]++
	}
	var out []*Decl
	for _, d := range prog.Decls {
		count[d]--
		if count[d] == 0 {
			out = append(out, d)
		}
	}
	prog.Decls = out
}
