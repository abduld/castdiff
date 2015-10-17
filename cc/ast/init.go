package ast

import "encoding/json"

// Init is an initializer expression.
type Init struct {
	SyntaxInfo
	Kind string
	Id     int
	Prefix []*Prefix // list of prefixes
	Expr   Expr      // Expr
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
	ret := ""
	for _, elem := range x.Prefix {
		ret += elem.String() + "\n"
	}
	return ret
}

func (x *Init) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "Init"
	}
	return json.Marshal(*x)

}
func (x *Init) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
