package ast

import "encoding/json"

type StringLiteral struct {
	SyntaxInfo
	Kind  string `json:"kind"`
	Id    int    `json:"id"`
	Value string `json:"value"`
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

func (x *StringLiteral) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "StringLiteral"
	}
	return json.Marshal(*x)

}
func (x *StringLiteral) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
