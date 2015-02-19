package ast

import "fmt"

type Decl struct {
	SyntaxInfo
	Id      int
	Name    Syntax
	Type    *Type
	Storage Storage
	Init    *Init
	Body    *Stmt

	XOuter    *Decl
	CurFn     *Decl
	OuterType *Type
}

func (x *Decl) GetId() int {
	return x.Id
}

func (x *Decl) GetChildren() []Syntax {
	lst := []Syntax{}
	if x.Type != nil {
		lst = append(lst, x.Type)
	}
	if x.Name != nil {
		lst = append(lst, x.Name)
	}
	if x.Init != nil {
		lst = append(lst, x.Init)
	}
	if x.Body != nil {
		lst = append(lst, x.Body)
	}
	return lst
}

func (d *Decl) String() string {
	if d == nil {
		return "nil Decl"
	}
	if d.Init != nil {
		return fmt.Sprintf("Decl<%d>{%s, %s} = %s", d.Id, d.Name.String(), d.Type, d.Init)
	} else {
		return fmt.Sprintf("Decl<%d>{%s, %s}", d.Id, d.Name.String(), d.Type)
	}
}
