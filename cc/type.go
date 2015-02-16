// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"fmt"
	"strconv"
	"strings"
)

var printf = fmt.Printf

type Type struct {
	SyntaxInfo
	Id       int
	Kind     TypeKind
	Qual     TypeQual
	Base     *Type
	Tag      Keyword
	Decls    []*Decl
	Width    *Expr
	Name     SymbolLiteral
	TypeDecl *Decl
}

func (x *Type) GetId() int {
	return x.Id
}

func (x *Type) GetChildren() []Syntax {
	lst := []Syntax{}
	if x.Base != nil {
		lst = append(lst, x.Base)
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

type TypeKind int

const (
	_ TypeKind = iota
	Void
	Char
	Uchar
	Short
	Ushort
	Int
	Uint
	Long
	Ulong
	Longlong
	Ulonglong
	Float
	Double
	Enum
	Ptr
	Struct
	Union
	Array
	Func
	TypedefType
)

var typeKindString = []string{
	Void:        "void",
	Char:        "char",
	Uchar:       "uchar",
	Short:       "short",
	Ushort:      "ushort",
	Int:         "int",
	Uint:        "uint",
	Long:        "long",
	Ulong:       "ulong",
	Longlong:    "longlong",
	Ulonglong:   "ulonglong",
	Float:       "float",
	Double:      "double",
	Ptr:         "pointer",
	Struct:      "struct",
	Union:       "union",
	Enum:        "enum",
	Array:       "array",
	Func:        "func",
	TypedefType: "<typedef>",
}

func (k TypeKind) String() string {
	if 0 <= int(k) && int(k) <= len(typeKindString) && typeKindString[k] != "" {
		return typeKindString[k]
	}
	return fmt.Sprintf("TypeKind(%d)", k)
}

type TypeQual int

const (
	Const TypeQual = 1 << iota
	Volatile
)

func (q TypeQual) String() string {
	s := ""
	if q&Const != 0 {
		s += "const "
	}
	if q&Volatile != 0 {
		s += "volatile "
	}
	if s == "" {
		return ""
	}
	return s[:len(s)-1]
}

type Storage int

const (
	Auto Storage = 1 << iota
	Static
	Extern
	Typedef
	Register
	Inline
)

func (c Storage) String() string {
	s := ""
	if c&Auto != 0 {
		s += "auto "
	}
	if c&Static != 0 {
		s += "static "
	}
	if c&Extern != 0 {
		s += "extern "
	}
	if c&Typedef != 0 {
		s += "typedef "
	}
	if c&Register != 0 {
		s += "register "
	}
	if c&Inline != 0 {
		s += "inline "
	}
	if s == "" {
		return ""
	}
	return s[:len(s)-1]
}

var (
	CharType      = newType(Char)
	UcharType     = newType(Uchar)
	ShortType     = newType(Short)
	UshortType    = newType(Ushort)
	IntType       = newType(Int)
	UintType      = newType(Uint)
	LongType      = newType(Long)
	UlongType     = newType(Ulong)
	LonglongType  = newType(Longlong)
	UlonglongType = newType(Ulonglong)
	FloatType     = newType(Float)
	DoubleType    = newType(Double)
	VoidType      = newType(Void)
	BoolType      = &Type{Kind: TypedefType, Name: "bool", Base: IntType}
)

type typeOp int

const (
	tChar typeOp = 1 << iota
	tShort
	tInt
	tLong
	tSigned
	tUnsigned
	tFloat
	tDouble
	tVoid
	tLonglong
)

var builtinTypes = map[typeOp]*Type{
	tChar:                     CharType,
	tChar | tSigned:           CharType,
	tChar | tUnsigned:         UcharType,
	tShort:                    ShortType,
	tShort | tSigned:          ShortType,
	tShort | tUnsigned:        UshortType,
	tShort | tInt:             ShortType,
	tShort | tSigned | tInt:   ShortType,
	tShort | tUnsigned | tInt: UshortType,
	tInt:                         IntType,
	tInt | tSigned:               IntType,
	tInt | tUnsigned:             UintType,
	tLong:                        LongType,
	tLong | tSigned:              LongType,
	tLong | tUnsigned:            UlongType,
	tLong | tInt:                 LongType,
	tLong | tSigned | tInt:       LongType,
	tLong | tUnsigned | tInt:     UlongType,
	tLonglong:                    LonglongType,
	tLonglong | tSigned:          LonglongType,
	tLonglong | tUnsigned:        UlonglongType,
	tLonglong | tInt:             LonglongType,
	tLonglong | tSigned | tInt:   LonglongType,
	tLonglong | tUnsigned | tInt: UlonglongType,
	tFloat:  FloatType,
	tDouble: DoubleType,
	tVoid:   VoidType,
}

func splitTypeWords(ws []string) (c Storage, q TypeQual, ty *Type) {
	// Could check for doubled words in general,
	// like const const, but no one cares.
	var t typeOp
	var ts []string
	for _, w := range ws {
		switch w {
		case "const":
			q |= Const
		case "volatile":
			q |= Volatile
		case "auto":
			c |= Auto
		case "static":
			c |= Static
		case "extern":
			c |= Extern
		case "typedef":
			c |= Typedef
		case "register":
			c |= Register
		case "inline":
			c |= Inline
		case "char":
			t |= tChar
			ts = append(ts, w)
		case "short":
			t |= tShort
			ts = append(ts, w)
		case "int":
			t |= tInt
			ts = append(ts, w)
		case "long":
			if t&tLong != 0 {
				t ^= tLonglong | tLong
			} else {
				t |= tLong
			}
			ts = append(ts, w)
		case "signed":
			t |= tSigned
			ts = append(ts, w)
		case "unsigned":
			t |= tUnsigned
			ts = append(ts, w)
		case "float":
			t |= tFloat
			ts = append(ts, w)
		case "double":
			t |= tDouble
			ts = append(ts, w)
		case "void":
			t |= tVoid
			ts = append(ts, w)
		}
	}

	if t == 0 {
		t |= tInt
	}

	ty = builtinTypes[t]
	if ty == nil {
		fmt.Printf("unsupported type %q\n", strings.Join(ts, " "))
	}

	return c, q, builtinTypes[t]
}

func newType(k TypeKind) *Type {
	return &Type{Kind: k}
}

func (t *Type) String() string {
	if t == nil {
		return "<nil>"
	}
	switch t.Kind {
	default:
		return t.Kind.String() + "<" + strconv.Itoa(t.Id) + ">"
	case TypedefType:
		if t.Name == "" {
			return "missing_typedef_name"
		}
		return t.Name
	case Ptr:
		return t.Base.String() + "*"
	case Struct, Union, Enum:
		if t.Tag == "" {
			return t.Kind.String()
		}
		return t.Kind.String() + " " + t.Tag
	case Array:
		return t.Base.String() + "[]"
	case Func:
		s := "func("
		for i, d := range t.Decls {
			if i > 0 {
				s += ", "
			}
			s += d.Name + " " + d.Type.String()
		}
		if t.Base == t {
			s += ") SELF"
		} else {
			s += ") " + t.Base.String()
		}
		return s
	}
}

type Decl struct {
	SyntaxInfo
	Id      int
	Name    SymbolLiteral
	Type    *Type
	Storage Storage
	Init    *Init
	Body    *Stmt

	XOuter    *Decl
	CurFn     *Decl
	OuterType *Type
}

func (x *Decl) GetId() int {
	return x.Id
}

func (x *Decl) GetChildren() []Syntax {
	lst := []Syntax{}
	if x.Type != nil {
		lst = append(lst, x.Type)
	}
	if x.Init != nil {
		lst = append(lst, x.Init)
	}
	if x.Body != nil {
		lst = append(lst, x.Body)
	}
	return lst
}

func (d *Decl) String() string {
	if d == nil {
		return "nil Decl"
	}
	if d.Init != nil {
		return fmt.Sprintf("Decl<%d>{%s, %s} = %s", d.Id, d.Name, d.Type, d.Init)
	} else {
		return fmt.Sprintf("Decl<%d>{%s, %s}", d.Id, d.Name, d.Type)
	}
}
