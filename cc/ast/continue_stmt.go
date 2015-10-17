package ast

import "encoding/json"

type ContinueStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
}

func (x *ContinueStmt) GetId() int {
	return x.Id
}

func (x *ContinueStmt) String() string {
	return "continue"
}

func (x *ContinueStmt) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *ContinueStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "ContinueStmt"
	}
	return json.Marshal(*x)

}
func (x *ContinueStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}

func (x *ContinueStmt) IsStmt() bool {
	return true
}