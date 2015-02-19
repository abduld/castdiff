// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"fmt"
)

type Expr interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

type ExprOp int

const (
	_          ExprOp = iota
	Add               // Left + Right
	AddEq             // Left += Right
	Addr              // &Left
	And               // Left & Right
	AndAnd            // Left && Right
	AndEq             // Left &= Right
	Arrow             // Left->Text
	Call              // Left(List)
	CUDACall          // Left(LaunchParams, List)
	Cast              // (Type)Left
	CastInit          // (Type){Init}
	Comma             // x, y, z; List = {x, y, z}
	Cond              // x ? y : z; List = {x, y, z}
	Div               // Left / Right
	DivEq             // Left /= Right
	Dot               // Left.Name
	Eq                // Left = Right
	EqEq              // Left == Right
	Gt                // Left > Right
	GtEq              // Left >= Right
	Index             // Left[Right]
	Indir             // *Left
	Lsh               // Left << Right
	LshEq             // Left <<= Right
	LcuBrk            // Left <<< Right
	Lt                // Left < Right
	LtEq              // Left <= Right
	RcuBrk            // Left >>> Right
	Minus             // -Left
	Mod               // Left % Right
	ModEq             // Left %= Right
	Mul               // Left * Right
	MulEq             // Left *= Right
	Name              // Text (function, variable, or enum name)
	Not               // !Left
	NotEq             // Left != Right
	Number            // Text (numeric or chraracter constant)
	Literal           // Text (numeric or chraracter constant)
	Offsetof          // offsetof(Type, Left)
	Or                // Left | Right
	OrEq              // Left |= Right
	OrOr              // Left || Right
	Paren             // (Left)
	Plus              //  +Left
	PostDec           // Left--
	PostInc           // Left++
	PreDec            // --Left
	PreInc            // ++Left
	Rsh               // Left >> Right
	RshEq             // Left >>= Right
	SizeofExpr        // sizeof(Left)
	SizeofType        // sizeof(Type)
	String            // Text (quoted string literal)
	Sub               // Left - Right
	SubEq             // Left -= Right
	Twid              // ~Left
	VaArg             // va_arg(Left, Type)
	Xor               // Left ^ Right
	XorEq             // Left ^= Right
	LCuBrk
	RCuBrk
)

var exprOpString = []string{
	Add:        "Add",
	AddEq:      "AddEq",
	Addr:       "Addr",
	And:        "And",
	AndAnd:     "AndAnd",
	AndEq:      "AndEq",
	Arrow:      "Arrow",
	Call:       "Call",
	Cast:       "Cast",
	CastInit:   "CastInit",
	Comma:      "Comma",
	Cond:       "Cond",
	Div:        "Div",
	DivEq:      "DivEq",
	Dot:        "Dot",
	Eq:         "Eq",
	EqEq:       "EqEq",
	Gt:         "Gt",
	GtEq:       "GtEq",
	Index:      "Index",
	Indir:      "Indir",
	Lsh:        "Lsh",
	LshEq:      "LshEq",
	Lt:         "Lt",
	LtEq:       "LtEq",
	Minus:      "Minus",
	Mod:        "Mod",
	ModEq:      "ModEq",
	Mul:        "Mul",
	MulEq:      "MulEq",
	Name:       "Name",
	Not:        "Not",
	NotEq:      "NotEq",
	Number:     "Number",
	Literal:    "Literal",
	Offsetof:   "Offsetof",
	Or:         "Or",
	OrEq:       "OrEq",
	OrOr:       "OrOr",
	Paren:      "Paren",
	Plus:       "Plus",
	PostDec:    "PostDec",
	PostInc:    "PostInc",
	PreDec:     "PreDec",
	PreInc:     "PreInc",
	Rsh:        "Rsh",
	RshEq:      "RshEq",
	SizeofExpr: "SizeofExpr",
	SizeofType: "SizeofType",
	String:     "String",
	Sub:        "Sub",
	SubEq:      "SubEq",
	Twid:       "Twid",
	VaArg:      "VaArg",
	Xor:        "Xor",
	XorEq:      "XorEq",
	LCuBrk:     "LCuBrk",
	RCuBrk:     "RCuBrk",
}

func (op ExprOp) String() string {
	if 0 <= int(op) && int(op) <= len(exprOpString) {
		return exprOpString[op]
	}
	return fmt.Sprintf("ExprOp(%d)", op)
}
