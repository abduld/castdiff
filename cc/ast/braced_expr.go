package ast

import "encoding/json"

type BracedExpr struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Expr Expr   `json:"expr"`
}

func (x *BracedExpr) GetId() int {
	return x.Id
}

func (x *BracedExpr) String() string {
	if x.Expr != nil {
		return "{ " + x.Expr.String() + " }"
	}
	return "{ }"
}

func (x *BracedExpr) GetChildren() []Syntax {
	return []Syntax{x.Expr}
}

func (x *BracedExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "BracedExpr"
	}
	return json.Marshal(*x)

}
func (x *BracedExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
