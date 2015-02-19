package ast

// Init is an initializer expression.
type Init struct {
	SyntaxInfo
	Id     int
	Prefix []*Prefix // list of prefixes
	Expr   *Expr     // Expr
	Braced []*Init   // {Braced}

	XType *Type // derived type
}

func (x *Init) GetId() int {
	return x.Id
}

func (x *Init) GetChildren() []Syntax {
	lst := []Syntax{}
	for _, elem := range x.Prefix {
		lst = append(lst, elem)
	}
	if x.Expr != nil {
		lst = append(lst, x.Expr)
	}
	for _, elem := range x.Braced {
		lst = append(lst, elem)
	}
	return lst
}

func (x *Init) String() string {
	var p Printer
	p.hideComments = true
	p.printInit(x)
	return p.String()
}
