package ast

import "encoding/json"

type ArrowOp int

const (
	_   ArrowOp = iota + 0x510
	Arrow       // Left.Name
)

type ArrowExpr struct {
	SyntaxInfo
	Kind   string `json:"kind"`
	Id     int    `json:"id"`
	Ptr Expr   `json:"ptr"`
	Member Expr   `json:"member"`
}

func (x *ArrowExpr) GetId() int {
	return x.Id
}

func (x *ArrowExpr) String() string {
	return x.Ptr.String() + "->" + x.Member.String()
}

func (x *ArrowExpr) GetChildren() []Syntax {
	return []Syntax{x.Ptr, x.Member}
}

func (x *ArrowExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "ArrowExpr"
	}
	return json.Marshal(*x)

}
func (x *ArrowExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
