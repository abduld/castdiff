package ast

import (
	"encoding/json"
	"strconv"
)

type IntegerLiteral struct {
	SyntaxInfo
	Id    int
	Value int
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
	return json.Marshal(*x)

}
func (x *IntegerLiteral) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
