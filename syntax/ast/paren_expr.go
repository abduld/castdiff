package ast

import "encoding/json"

type ParenExpr struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Expr Expr   `json:"expr"`
}

func (x *ParenExpr) GetId() int {
	return x.Id
}

func (x *ParenExpr) String() string {
	return "( " + x.Expr.String() + " )"
}

func (x *ParenExpr) GetChildren() []Syntax {
	return []Syntax{x.Expr}
}

func (x *ParenExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "ParenExpr"
	}
	return json.Marshal(*x)

}
func (x *ParenExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
