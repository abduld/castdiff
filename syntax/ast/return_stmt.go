package ast

import "encoding/json"

type ReturnStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
}

func (x *ReturnStmt) GetId() int {
	return x.Id
}

func (x *ReturnStmt) String() string {
	return "return"
}

func (x *ReturnStmt) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *ReturnStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "ReturnStmt"
	}
	return json.Marshal(*x)

}
func (x *ReturnStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
