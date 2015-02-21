package ast

import "fmt"

type AssignOp int

const (
	_     AssignOp = iota
	AddEq          // Left += Right
	AndEq          // Left &= Right
	DivEq          // Left /= Right
	Eq             // Left = Right
	LshEq          // Left <<= Right
	ModEq          // Left %= Right
	MulEq          // Left *= Right
	OrEq           // Left |= Right
	RshEq          // Left >>= Right
	SubEq          // Left -= Right
	XorEq          // Left ^= Right
)

var assignOpString = []string{

	AddEq: "AddEq",
	AndEq: "AndEq",
	DivEq: "DivEq",
	Eq:    "Eq",
	LshEq: "LshEq",
	ModEq: "ModEq",
	MulEq: "MulEq",
	OrEq:  "OrEq",
	RshEq: "RshEq",
	SubEq: "SubEq",
	XorEq: "XorEq",
}

var assignOpInfix = []string{
	AddEq: "+=",
	AndEq: "&=",
	DivEq: "/=",
	Eq:    "=",
	LshEq: "<<=",
	ModEq: "%=",
	MulEq: "*=",
	OrEq:  "|=",
	RshEq: ">>=",
	SubEq: "-=",
	XorEq: "^=",
}

func (op AssignOp) Name() string {
	if 0 <= int(op) && int(op) <= len(assignOpString) {
		return assignOpString[op]
	}
	return fmt.Sprintf("AssignOp(%d)", op)
}

func (op AssignOp) String() string {
	if 0 <= int(op) && int(op) <= len(assignOpInfix) {
		return assignOpInfix[op]
	}
	return fmt.Sprintf("AssignOp(%d)", op)
}
