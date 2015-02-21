package ast

import (
	"encoding/json"
	"strings"
)

type CommaExpr struct {
	SyntaxInfo
	Kind  string `json:"kind"`
	Id    int    `json:"id"`
	Exprs []Expr `json:"exprs"`
}

func (x *CommaExpr) GetId() int {
	return x.Id
}

func (x *CommaExpr) GetChildren() []Syntax {
	sargs := make([]Syntax, len(x.Exprs))
	for ii, arg := range x.Exprs {
		sargs[ii] = arg
	}
	return sargs
}

func (x *CommaExpr) String() string {
	sargs := make([]string, len(x.Exprs))
	for ii, arg := range x.Exprs {
		sargs[ii] = arg.String()
	}
	return strings.Join(sargs, ",")
}

func (x *CommaExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "CommaExpr"
	}
	return json.Marshal(*x)

}

func (x *CommaExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
