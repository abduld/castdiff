package ast

import "fmt"

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
