package ast

import "encoding/json"

type ExprStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Expr   Expr    `json:"expr"`
}

func (x *ExprStmt) GetId() int {
	return x.Id
}

func (x *ExprStmt) String() string {
	return x.Expr.String() + ";"
}

func (x *ExprStmt) GetChildren() []Syntax {
	return []Syntax{x.Expr}
}

func (x *ExprStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "ExprStmt"
	}
	return json.Marshal(*x)

}
func (x *ExprStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
func (x *ExprStmt) IsStmt() bool {
	return true
}

