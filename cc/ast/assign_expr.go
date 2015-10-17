package ast

import "encoding/json"

type AssignExpr struct {
	SyntaxInfo
	Kind  string   `json:"kind"`
	Id    int      `json:"id"`
	Op    AssignOp `json:"op"`
	Left  Expr     `json:"left"`
	Right Expr     `json:"right"`
}

func (x *AssignExpr) GetId() int {
	return x.Id
}

func (x *AssignExpr) GetChildren() []Syntax {
	return []Syntax{x.Left, x.Right}
}

func (x *AssignExpr) String() string {
	str := x.Right.String()
	str += " " + x.Op.String() + " "
	str += x.Left.String()
	return str
}

func (x *AssignExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "AssignExpr"
	}
	return json.Marshal(*x)

}

func (x *AssignExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
