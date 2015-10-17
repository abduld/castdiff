package ast

import "encoding/json"

type DotExpr struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Struct Expr   `json:"struct"`
	Member Expr   `json:"member"`
}

func (x *DotExpr) GetId() int {
	return x.Id
}

func (x *DotExpr) String() string {
	return x.Struct.String() + "." + x.Member.String()
}

func (x *DotExpr) GetChildren() []Syntax {
	return []Syntax{x.Struct, x.Member}
}

func (x *DotExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "DotExpr"
	}
	return json.Marshal(*x)

}
func (x *DotExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
