package ast

import "encoding/json"

type EmptyLiteral struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
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
	if x != nil {
		x.Kind = "EmptyLiteral"
	}
	return json.Marshal(*x)

}
func (x *EmptyLiteral) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
