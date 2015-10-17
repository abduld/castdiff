package ast

import "encoding/json"

type LanguageKeyword struct {
	SyntaxInfo
	Kind  string `json:"kind"`
	Id    int    `json:"id"`
	Value string    `json:"value"`
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

func (x *LanguageKeyword) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "LanguageKeyword"
	}
	return json.Marshal(*x)

}
func (x *LanguageKeyword) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
