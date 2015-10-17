package ast

type SyntaxOp int

const (
	_       SyntaxOp = iota + 0x30
	Comma            // x, y, z; List = {x, y, z}
	Name             // Text (function, variable, or enum name)
	String           // Text (quoted string literal)
	Number           // Text (numeric or chraracter constant)
	Literal          // Text (numeric or chraracter constant)
	VaArg            // Text (numeric or chraracter constant)
	LcuBrk           // left cuda braket
	RcuBrk           // right cuda bracket
)
