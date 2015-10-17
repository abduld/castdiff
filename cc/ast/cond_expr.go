package ast

import "encoding/json"

type CondOp int

const (
	_    CondOp = iota + 0x400
	Cond        // x ? y : z; List = {x, y, z}
)

type CondExpr struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Cond Expr   `json:"cond"`
	Then Expr   `json:"then"`
	Else Expr   `json:"else"`
}

func (x *CondExpr) GetId() int {
	return x.Id
}

func (x *CondExpr) String() string {
	return x.Cond.String() + " ? " + x.Then.String() + " : " + x.Else.String()
}

func (x *CondExpr) GetChildren() []Syntax {
	return []Syntax{x.Cond, x.Then, x.Else}
}

func (x *CondExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "CondExpr"
	}
	return json.Marshal(*x)

}
func (x *CondExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
