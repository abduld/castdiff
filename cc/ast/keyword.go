package ast

type LanguageKeyword struct {
	SyntaxInfo
	Id    int
	Value string
}

func (x *LanguageKeyword) GetId() int {
	return x.Id
}

func (x *LanguageKeyword) String() string {
	if x == nil {
		return ""
	} else {
		return x.Value
	}
}

func (x *LanguageKeyword) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *LanguageKeyword) ToSymbolLiteral() *SymbolLiteral {
	return &SymbolLiteral{
		SyntaxInfo: x.SyntaxInfo,
		Value:      x.Value,
		Id:         x.Id,
	}
}
