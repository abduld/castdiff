package ast

import (
	"encoding/json"
	"strconv"
)

type IntegerLiteral struct {
	SyntaxInfo
	Kind  string `json:"kind"`
	Id    int    `json:"id"`
	Value int    `json:"value"`
}

func (x *IntegerLiteral) GetId() int {
	return x.Id
}

func (x *IntegerLiteral) String() string {
	if x == nil {
		return ""
	} else {
		return strconv.Itoa(x.Value)
	}
}

func (x *IntegerLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *IntegerLiteral) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "IntegerLiteral"
	}
	return json.Marshal(*x)

}
func (x *IntegerLiteral) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
