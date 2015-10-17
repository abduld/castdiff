package ast

import "encoding/json"

type ArrayExpr struct {
	SyntaxInfo
	Kind  string     `json:"kind"`
	Id    int        `json:"id"`
	Array Expr       `json:"array"`
	Index *TupleExpr `json:"index"`
}

func (x *ArrayExpr) GetId() int {
	return x.Id
}

func (x *ArrayExpr) GetChildren() []Syntax {
	children := make([]Syntax, len(x.Index.Args)+1)
	children[0] = x.Array
	for ii, arg := range x.Index.Args {
		children[ii+1] = arg
	}
	return children
}

func (x *ArrayExpr) String() string {
	str := x.Array.String()
	str += x.Index.String()
	return str
}

func (x *ArrayExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "ArrayExpr"
	}
	return json.Marshal(*x)

}

func (x *ArrayExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
