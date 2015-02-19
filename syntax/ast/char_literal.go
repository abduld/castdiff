package ast

import "encoding/json"

type CharLiteral struct {
	SyntaxInfo
	Id    int
	Value byte
}

func (x *CharLiteral) GetId() int {
	return x.Id
}

func (x *CharLiteral) String() string {
	return "'" + string([]byte{x.Value}) + "'"
}

func (x *CharLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *CharLiteral) MarshalJSON() ([]byte, error) {
	return json.Marshal(*x)

}
func (x *CharLiteral) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
