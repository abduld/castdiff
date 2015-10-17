package ast

import "fmt"

type UnaryOp int

const (
	_       UnaryOp = iota + 0x200
	Addr            // &Left
	Indir           // *Left
	Paren           // (Left)
	Plus            //  +Left
	PostDec         // Left--
	PostInc         // Left++
	PreDec          // --Left
	PreInc          // ++Left
	Twid            // ~Left
	Minus           // -Left
	Not             // !Left
)

var unaryOpString = []string{
	Addr:    "Addr",
	Indir:   "Indir",
	Paren:   "Paren",
	Plus:    "Plus",
	PostDec: "PostDec",
	PostInc: "PostInc",
	PreDec:  "PreDec",
	PreInc:  "PreInc",
	Twid:    "Twid",
	Minus:   "Minus",
	Not:     "Not",
}

var unaryOpPrefix = []string{
	Addr:    "&",
	Indir:   "*",
	Paren:   "(",
	Plus:    "+",
	PostDec: "",
	PostInc: "",
	PreDec:  "--",
	PreInc:  "++",
	Twid:    "~",
	Minus:   "-",
	Not:     "!",
}

var unaryOpPostfix = []string{
	Addr:    "",
	Indir:   "",
	Paren:   ")",
	Plus:    "",
	PostDec: "++",
	PostInc: "--",
	PreDec:  "",
	PreInc:  "",
	Twid:    "",
	Minus:   "",
	Not:     "",
}

func (op UnaryOp) Name() string {
	if 0 <= int(op) && int(op) <= len(unaryOpString) {
		return unaryOpString[op]
	}
	return fmt.Sprintf("UnaryOp(%d)", op)
}

func (op UnaryOp) Prefix() string {
	if 0 <= int(op) && int(op) <= len(unaryOpPrefix) {
		return unaryOpPrefix[op]
	}
	return ""
}

func (op UnaryOp) Postfix() string {
	if 0 <= int(op) && int(op) <= len(unaryOpPrefix) {
		return unaryOpPrefix[op]
	}
	return ""
}

func (op UnaryOp) String() string {
	if 0 <= int(op) && int(op) <= len(unaryOpPrefix) {
		return unaryOpPrefix[op] + unaryOpPostfix[op]
	}
	return fmt.Sprintf("UnaryOp(%d)", op)
}
