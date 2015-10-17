package ast

import "encoding/json"

type CastOp int

const (
	_        CastOp = iota + 0x10000
	Cast            // (Type)Left
	CastInit        // (Type){Init}
)

type CastExpr struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Type *Type  `json:"type"`
	Expr Expr   `json:"expr"`
}

func (x *CastExpr) GetId() int {
	return x.Id
}

func (x *CastExpr) GetChildren() []Syntax {
	return []Syntax{x.Type, x.Expr}
}

func (x *CastExpr) String() string {
	return "(" + x.Type.String() + ") " + x.Expr.String()
}

func (x *CastExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "CastExpr"
	}
	return json.Marshal(*x)

}

func (x *CastExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
