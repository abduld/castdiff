package ast

import (
	"encoding/json"
	"strings"
)

type CallExpr struct {
	SyntaxInfo
	Kind   string `json:"kind"`
	Id     int    `json:"id"`
	Callee Expr   `json:"callee"`
	Args   []Expr `json:"args"`
}

func (x *CallExpr) GetId() int {
	return x.Id
}

func (x *CallExpr) String() string {
	str := x.Callee.String()
	sargs := make([]string, len(x.Args))
	for ii, arg := range x.Args {
		sargs[ii] = arg.String()
	}
	str += "(" + strings.Join(sargs, ",") + ")"
	return str
}
func (x *CallExpr) GetChildren() []Syntax {
	children := make([]Syntax, len(x.Args)+1)
	children[0] = x.Callee
	for ii, arg := range x.Args {
		children[ii+1] = arg
	}
	return children
}

func (x *CallExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "CallExpr"
	}
	return json.Marshal(*x)

}
func (x *CallExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
