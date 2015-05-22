package ast

import (
	"encoding/json"
	"strconv"
)

type RealLiteral struct {
	SyntaxInfo
	Id    int
	Value float64
}

func (x *RealLiteral) GetId() int {
	return x.Id
}

func (x *RealLiteral) String() string {
	return strconv.FormatFloat(x.Value, 'f', 6, 64)
}

func (x *RealLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *RealLiteral) MarshalJSON() ([]byte, error) {
	return json.Marshal(*x)

}
func (x *RealLiteral) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
