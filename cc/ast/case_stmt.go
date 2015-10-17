package ast

import "encoding/json"

type CaseStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Expr Expr   `json:"expr"`
}

func (x *CaseStmt) GetId() int {
	return x.Id
}

func (x *CaseStmt) String() string {
	return "case " + x.Expr.String() + ":"
}

func (x *CaseStmt) GetChildren() []Syntax {
	return []Syntax{x.Expr}
}

func (x *CaseStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "CaseStmt"
	}
	return json.Marshal(*x)

}
func (x *CaseStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}

func (x *CaseStmt) IsStmt() bool {
	return true
}
