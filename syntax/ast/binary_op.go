package ast

import "fmt"

type BinaryOp int

const (
	_      BinaryOp = iota
	Add             // Left + Right
	And             // Left & Right
	AndAnd          // Left && Right
	Arrow           // Left->Right
	Div             // Left / Right
	EqEq            // Left == Right
	Gt              // Left > Right
	GtEq            // Left >= Right
	Lsh             // Left << Right
	Lt              // Left < Right
	LtEq            // Left <= Right
	Mod             // Left % Right
	Mul             // Left * Right
	NotEq           // Left != Right
	Or              // Left | Right
	OrOr            // Left || Right
	Rsh             // Left >> Right
	Sub             // Left - Right
	Xor             // Left ^ Right
)

var binaryOpString = []string{
	Add:    "Add",
	And:    "And",
	AndAnd: "AndAnd",
	Arrow:  "Arrow",
	Div:    "Div",
	EqEq:   "EqEq",
	Gt:     "Gt",
	GtEq:   "GtEq",
	Lsh:    "Lsh",
	Lt:     "Lt",
	LtEq:   "Lt",
	Mod:    "Mod",
	Mul:    "Mul",
	NotEq:  "NotEq",
	Or:     "Or",
	OrOr:   "OrOr",
	Rsh:    "Rsh",
	Sub:    "Sub",
	Xor:    "Xor",
}

var binaryOpInfix = []string{
	Add:    "+",
	And:    "&",
	AndAnd: "&&",
	Arrow:  "->",
	Div:    "/",
	EqEq:   "==",
	Gt:     ">",
	GtEq:   ">=",
	Lsh:    "<<",
	Lt:     "<",
	LtEq:   "<",
	Mod:    "%",
	Mul:    "*",
	NotEq:  "!=",
	Or:     "|",
	OrOr:   "||",
	Rsh:    ">>",
	Sub:    "-",
	Xor:    "^",
}

func (op BinaryOp) Name() string {
	if 0 <= int(op) && int(op) <= len(binaryOpString) {
		return binaryOpString[op]
	}
	return fmt.Sprintf("BinaryOp(%d)", op)
}

func (op BinaryOp) String() string {
	if 0 <= int(op) && int(op) <= len(binaryOpInfix) {
		return binaryOpInfix[op]
	}
	return fmt.Sprintf("BinaryOp(%d)", op)
}
