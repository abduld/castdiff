package ast

type StringLiteral struct {
	SyntaxInfo
	Id    int
	Value string
}

func (x *StringLiteral) GetId() int {
	return x.Id
}

func (x *StringLiteral) String() string {
	if x == nil {
		return `""`
	} else {
		return x.Value
	}
}

func (x *StringLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *StringLiteral) ToSymbolLiteral() *SymbolLiteral {
	return &SymbolLiteral{
		SyntaxInfo: x.SyntaxInfo,
		Value:      x.Value,
		Id:         x.Id,
	}
}
