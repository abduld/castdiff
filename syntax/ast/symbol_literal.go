package ast

import "encoding/json"

type SymbolLiteral struct {
	SyntaxInfo
	Kind  string `json:"kind"`
	Id    int    `json:"id"`
	Value string `json:"value"`
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

func (x *SymbolLiteral) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "SymbolLiteral"
	}
	return json.Marshal(*x)

}
func (x *SymbolLiteral) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
