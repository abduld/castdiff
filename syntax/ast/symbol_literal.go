package ast

type SymbolLiteral struct {
	SyntaxInfo
	Id    int
	Value string
}

func (x *SymbolLiteral) GetId() int {
	return x.Id
}

func (x *SymbolLiteral) String() string {
	if x == nil {
		return ""
	} else {
		return x.Value
	}
}

func (x *SymbolLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *SymbolLiteral) ToSymbolLiteral() *SymbolLiteral {
	return x
}
