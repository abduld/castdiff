package ast

// A Syntax represents any syntax element.
type Syntax interface {

	// GetSpan returns the start and end position of the syntax,
	// excluding leading or trailing comments.
	// The use of a Get prefix is non-standard but avoids a conflict
	// with the field named Span in most implementations.
	GetSpan() Span

	// GetComments returns the comments attached to the syntax.
	// This method would normally be named 'Comments' but that
	// would interfere with embedding a type of the same name.
	// The use of a Get prefix is non-standard but avoids a conflict
	// with the field named Comments in most implementations.
	GetComments() *Comments

	GetId() int

	GetChildren() []Syntax

	String() string
}

// SyntaxInfo contains metadata about a piece of syntax.
type SyntaxInfo struct {
	Span     Span     `json:"span"` // location of syntax in input
	Comments Comments `json:"comments,omitempty"`
}

func (s *SyntaxInfo) GetSpan() Span {
	return s.Span
}

func (s *SyntaxInfo) String() string {
	return ""
}

func (s *SyntaxInfo) GetComments() *Comments {
	return &s.Comments
}
