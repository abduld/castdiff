// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

type Stmt struct {
	SyntaxInfo
	Id     int
	Op     StmtOp
	Pre    *Expr
	Expr   *Expr
	Post   *Expr
	Decl   *Decl
	Body   *Stmt
	Else   *Stmt
	Block  []*Stmt
	Labels []*Label
	Text   string
	Type   *Type
}

func (x *Stmt) GetId() int {
	return x.Id
}

func (x *Stmt) GetChildren() []Syntax {
	lst := []Syntax{}
	if len(x.Labels) > 0 {
		for _, elem := range x.Labels {
			lst = append(lst, elem)
		}
	}
	switch x.Op {
	case Block:
		for _, elem := range x.Block {
			lst = append(lst, elem)
		}
	case Do:
		lst = append(lst, x.Body, x.Expr)
	case For:
		if x.Pre != nil {
			lst = append(lst, x.Pre)
		}
		if x.Expr != nil {
			lst = append(lst, x.Expr)
		}
		if x.Post != nil {
			lst = append(lst, x.Post)
		}
		if x.Body != nil {
			lst = append(lst, x.Body)
		}
	case If:
		lst = append(lst, x.Expr)
		lst = append(lst, x.Body)
		if x.Else != nil {
			lst = append(lst, x.Else)
		}
	case Return:
		if x.Expr != nil {
			lst = append(lst, x.Expr)
		}
	case StmtDecl:
		if x.Decl != nil {
			lst = append(lst, x.Decl)
		}
	case StmtExpr:
		if x.Expr != nil {
			lst = append(lst, x.Expr)
		}
	case Switch:
		if x.Expr != nil {
			lst = append(lst, x.Expr)
		}
		if x.Body != nil {
			lst = append(lst, x.Body)
		}
	case While:
		if x.Expr != nil {
			lst = append(lst, x.Expr)
		}
		if x.Body != nil {
			lst = append(lst, x.Body)
		}
	}
	return lst
}

type StmtOp int

const (
	_ StmtOp = iota
	StmtDecl
	StmtExpr
	Empty
	Block
	ARGBEGIN
	Break
	Continue
	Do
	For
	If
	Goto
	Return
	Switch
	While
)

type Label struct {
	SyntaxInfo
	Id   int
	Op   LabelOp
	Expr *Expr
	Name string
}

func (x *Label) GetId() int {
	return x.Id
}

func (x *Label) GetChildren() []Syntax {
	if x.Expr != nil {
		return []Syntax{x.Expr}
	} else {
		return []Syntax{}
	}
}

type LabelOp int

const (
	_ LabelOp = iota
	Case
	Default
	LabelName
)
