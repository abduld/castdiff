package ast

import "fmt"

type DeclStmt struct {
	SyntaxInfo
	Id      int           `json:"id"`
	Type    Type          `json:"type"`
	Storage Storage       `json:"storage"`
	Name    SymbolLiteral `json:"name"`
	Init    Expr          `json:"init"`
}

func (x *DeclStmt) GetId() int {
	return x.Id
}

func (x *DeclStmt) GetChildren() []Syntax {
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
	return lst
}

func (d *DeclStmt) String() string {
	if d == nil {
		return "nil DeclStmt"
	}
	if d.Init != nil {
		return fmt.Sprintf("DeclStmt<%d>{%s, %s} = %s", d.Id, d.Name.String(), d.Type, d.Init)
	} else {
		return fmt.Sprintf("DeclStmt<%d>{%s, %s}", d.Id, d.Name.String(), d.Type)
	}
}
