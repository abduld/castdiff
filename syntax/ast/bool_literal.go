package ast

import "encoding/json"

type BooleanLiteral struct {
	SyntaxInfo
	Id    int
	Value bool
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
	return json.Marshal(*x)

}
func (x *BooleanLiteral) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
