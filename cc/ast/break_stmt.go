package ast

import "encoding/json"

type BreakStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
}

func (x *BreakStmt) GetId() int {
	return x.Id
}

func (x *BreakStmt) String() string {
	return "break"
}

func (x *BreakStmt) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *BreakStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "BreakStmt"
	}
	return json.Marshal(*x)

}
func (x *BreakStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
