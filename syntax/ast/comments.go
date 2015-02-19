package ast

// Comments collects the comments associated with a syntax element.
type Comments struct {
	Before []Comment // whole-line comments before this syntax
	Suffix []Comment // end-of-line comments after this syntax

	// For top-level syntax elements only, After lists whole-line
	// comments following the syntax.
	After []Comment
}
type Comment struct {
	Span
	Text   string
	Suffix bool
}
