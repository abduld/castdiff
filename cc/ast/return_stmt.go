package ast

import "encoding/json"

type ReturnStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Value Expr `json:"value"`
}

func (x *ReturnStmt) GetId() int {
	return x.Id
}

func (x *ReturnStmt) String() string {
	return "return " + x.Value.String() + ";"
}

func (x *ReturnStmt) GetChildren() []Syntax {
	return []Syntax{x.Value}
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

func (x *ReturnStmt) IsStmt() bool {
	return true
}