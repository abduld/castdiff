package ast

import "fmt"

type TypeType int

const (
	_ TypeType = iota
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

var TypeTypeString = []string{
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

func (k TypeType) String() string {
	if 0 <= int(k) && int(k) <= len(TypeTypeString) && TypeTypeString[k] != "" {
		return TypeTypeString[k]
	}
	return fmt.Sprintf("TypeType(%d)", k)
}
