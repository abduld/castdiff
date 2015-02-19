package ast

import "encoding/json"

type EmptyLiteral struct {
	SyntaxInfo
	Id int
}

func (x *EmptyLiteral) GetId() int {
	return x.Id
}

func (x *EmptyLiteral) String() string {
	return ""
}

func (x *EmptyLiteral) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *EmptyLiteral) MarshalJSON() ([]byte, error) {
	return json.Marshal(*x)

}
func (x *EmptyLiteral) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
