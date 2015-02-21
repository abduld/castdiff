// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import "strconv"

type Type struct {
	SyntaxInfo
	Id    int      `json:"id"`
	Kind  TypeKind `json:"kind"`
	Qual  TypeQual `json:"qual"`
	Base  *Type    `json:"base"`
	Tag   Syntax   `json:"tag"`
	Decls []*Decl  `json:"decls"`
	Name  Syntax   `json:"name"`
}

func (x *Type) GetId() int {
	return x.Id
}

func (x *Type) GetChildren() []Syntax {
	lst := []Syntax{}
	if x.Base != nil {
		lst = append(lst, x.Base)
	}
	if x.Tag != nil {
		lst = append(lst, x.Tag)
	}
	if x.Name != nil {
		lst = append(lst, x.Name)
	}
	for _, elem := range x.Decls {
		lst = append(lst, elem)
	}
	return lst
}

func (x *Type) GetComments() *Comments {
	return nil
}

func (s *Type) GetSpan() Span {
	return Span{}
}

func (t *Type) String() string {
	if t == nil {
		return "<nil>"
	}
	switch t.Kind {
	default:
		return t.Kind.String() + "<" + strconv.Itoa(t.Id) + ">"
	case TypedefType:
		if t.Name.String() == "" {
			return "missing_typedef_name"
		}
		return t.Name.String()
	case Ptr:
		return t.Base.String() + "*"
	case Struct, Union, Enum:
		if t.Tag.String() == "" {
			return t.Kind.String()
		}
		return t.Kind.String() + " " + t.Tag.String()
	case Array:
		return t.Base.String() + "[]"
	case Func:
		s := "func("
		for i, d := range t.Decls {
			if i > 0 {
				s += ", "
			}
			s += d.Name.String() + " " + d.Type.String()
		}
		if t.Base == t {
			s += ") SELF"
		} else {
			s += ") " + t.Base.String()
		}
		return s
	}
}
