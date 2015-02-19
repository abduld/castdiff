package ast

import "encoding/json"

type BooleanLiteral struct {
	SyntaxInfo
	Kind  string `json:"kind"`
	Id    int    `json:"id"`
	Value bool   `json:"value"`
}

func (x *BooleanLiteral) GetId() int {
	return x.Id
}

func (x *BooleanLiteral) String() string {
	if x.Value == true {
		return "true"
	} else {
		return "false"
	}
}

func (x *BooleanLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *BooleanLiteral) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "BooleanLiteral"
	}
	return json.Marshal(*x)

}
func (x *BooleanLiteral) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
