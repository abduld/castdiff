package ast

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
	BoolType      = &Type{Type: TypedefType, Name: &SymbolLiteral{Value: "bool"}, Base: IntType}
)

func newType(k TypeType) *Type {
	return &Type{Type: k}
}
