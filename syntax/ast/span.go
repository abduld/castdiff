package ast

type Span struct {
	Start Position `json:"end"`
	End   Position `json:"start"`
}
