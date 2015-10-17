package ast

import "encoding/json"

type UnaryExpr struct {
	SyntaxInfo
	Kind string  `json:"kind"`
	Id   int     `json:"id"`
	Op   UnaryOp `json:"op"`
	Arg  Expr    `json:"arg"`
}

func (x *UnaryExpr) GetId() int {
	return x.Id
}

func (x *UnaryExpr) String() string {
	str := x.Op.Prefix()
	str += x.Arg.String()
	str += x.Op.Postfix()
	return str
}

func (x *UnaryExpr) GetChildren() []Syntax {
	return []Syntax{x.Arg}
}

func (x *UnaryExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "UnaryExpr"
	}
	return json.Marshal(*x)

}
func (x *UnaryExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
