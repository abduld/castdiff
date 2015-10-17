package ast

import "encoding/json"

type VaArgExpr struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Type Expr `json:"type"`
	Arg Expr   `json:"arg"`
}

func (x *VaArgExpr) GetId() int {
	return x.Id
}

func (x *VaArgExpr) String() string {
	return "..."
}

func (x *VaArgExpr) GetChildren() []Syntax {
	return []Syntax{x.Arg}
}

func (x *VaArgExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "VaArgExpr"
	}
	return json.Marshal(*x)

}
func (x *VaArgExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
