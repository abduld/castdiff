package ast

import (
	"encoding/json"
	"strings"
)

type CUDALaunchExpr struct {
	SyntaxInfo
	Kind   string `json:"kind"`
	Id     int    `json:"id"`
	Callee Expr   `json:"CUDALaunchee"`
	LaunchParams []Expr `json:"launch_params"`
	Args   []Expr `json:"args"`
}

func (x *CUDALaunchExpr) GetId() int {
	return x.Id
}

func (x *CUDALaunchExpr) String() string {
	str := x.Callee.String()
	largs := make([]string, len(x.LaunchParams))
	for ii, arg := range x.LaunchParams {
		largs[ii] = arg.String()
	}
	sargs := make([]string, len(x.Args))
	for ii, arg := range x.Args {
		sargs[ii] = arg.String()
	}
	str += "<<<" + strings.Join(largs, ",") + ">>>"
	str += "(" + strings.Join(sargs, ",") + ")"
	return str
}
func (x *CUDALaunchExpr) GetChildren() []Syntax {
	children := make([]Syntax, len(x.LaunchParams)+len(x.Args)+1)
	children[0] = x.Callee
	for ii, arg := range x.LaunchParams {
		children[ii+1] = arg
	}
	for ii, arg := range x.Args {
		children[ii+len(x.LaunchParams)+1] = arg
	}
	return children
}

func (x *CUDALaunchExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "CUDALaunchExpr"
	}
	return json.Marshal(*x)

}
func (x *CUDALaunchExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
