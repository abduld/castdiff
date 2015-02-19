package ast

type Expr interface {
	Syntax
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
