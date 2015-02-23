package ast

// Prefix is an initializer prefix.
type Prefix struct {
	Span  Span
	Id    int
	Dot   Syntax // .Dot =
	XDecl *DeclStmt  // for .Dot
	Index *Expr  // [Index] =
}

func (x *Prefix) GetId() int {
	return x.Id
}

func (x *Prefix) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *Prefix) GetComments() *Comments {
	return nil
}

func (s *Prefix) GetSpan() Span {
	return Span{}
}

func (x *Prefix) String() string {
	if x.Dot.String() != "" {
		return "." + x.Dot.String()
	} else {
		return "[" + x.Index.String() + "]"
	}
}
