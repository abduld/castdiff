package ast

import "encoding/json"

type EmptyStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
}

func (x *EmptyStmt) GetId() int {
	return x.Id
}

func (x *EmptyStmt) String() string {
	return ""
}

func (x *EmptyStmt) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *EmptyStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "EmptyStmt"
	}
	return json.Marshal(*x)

}
func (x *EmptyStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
func (x *EmptyStmt) IsStmt() bool {
	return true
}

