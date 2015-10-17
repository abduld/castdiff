package ast

import (
	"encoding/json"
	"strings"
)

type TupleKind int

const (
	_ TupleKind = iota
	None
	Bracket
	Parenthesis
	Brace
)

type TupleExpr struct {
	SyntaxInfo
	Kind     string    `json:"kind"`
	Id       int       `json:"id"`
	Brackets TupleKind `json:"Tupleee"`
	Args     []Expr    `json:"args"`
}

func (x *TupleExpr) GetId() int {
	return x.Id
}

func (x *TupleExpr) Prefix() string {
	switch x.Brackets {
	case None:
		return ""
	case Bracket:
		return "["
	case Brace:
		return "{"
	case Parenthesis:
		return "("
	}
	return "u{"
}

func (x *TupleExpr) Postfix() string {
	switch x.Brackets {
	case None:
		return ""
	case Bracket:
		return "]"
	case Brace:
		return "}"
	case Parenthesis:
		return ")"
	}
	return "}u"
}

func (x *TupleExpr) String() string {
	sargs := make([]string, len(x.Args))
	for ii, arg := range x.Args {
		sargs[ii] = arg.String()
	}
	return x.Prefix() + strings.Join(sargs, ",") + x.Postfix()
}

func (x *TupleExpr) GetChildren() []Syntax {
	children := make([]Syntax, len(x.Args))
	for ii, stmt := range x.Args {
		children[ii] = stmt
	}
	return children
}

func (x *TupleExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "TupleExpr"
	}
	return json.Marshal(*x)

}
func (x *TupleExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
