// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import (
	"encoding/json"
)

type Type struct {
	SyntaxInfo
	Id    int      `json:"id"`
	Kind  string   `json:"kind"`
	Type  TypeType `json:"type"`
	Qual  TypeQual `json:"qual"`
	Base  *Type    `json:"base"`
	Size  Expr     `json:"size"`
	Tag   Syntax   `json:"tag"`
	Stmts []Stmt   `json:"stmts"`
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
	if x.Size != nil {
		lst = append(lst, x.Size)
	}
	for _, elem := range x.Stmts {
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
	switch t.Type {
	default:
		return t.Type.String()
	case TypedefType:
		if t.Name.String() == "" {
			return "missing_typedef_name"
		}
		return t.Name.String()
	case Ptr:
		return t.Base.String() + "*"
	case Struct, Union, Enum:
		if t.Tag.String() == "" {
			return t.Type.String()
		}
		return t.Type.String() + " " + t.Tag.String()
	case Array:
		if t.Size != nil {
			return t.Base.String() + "[" + t.Size.String() + "]"
		}
		return t.Base.String() + "[]"
	case Func:
		s := "func("
		for i, d := range t.Stmts {
			if i > 0 {
				s += ", "
			}

			switch d := d.(type) {
			default:
				continue
			case *DeclStmt:
				s += d.Name.String() + " " + d.Type.String()
			}
		}
		if t.Base == t {
			s += ") SELF"
		} else {
			s += ") " + t.Base.String()
		}
		return s
	}
}

func (x *Type) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "Type"
	}
	return json.Marshal(*x)

}
func (x *Type) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
