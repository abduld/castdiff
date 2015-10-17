package ast

import (
	"encoding/json"
	"fmt"
)

type DeclExpr struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id      int            `json:"id"`
	Type    *Type          `json:"type"`
	Storage Storage        `json:"storage"`
	Name    *SymbolLiteral `json:"name"`
	Init    Expr           `json:"init"`
}

func (x *DeclExpr) GetId() int {
	return x.Id
}

func (x *DeclExpr) GetChildren() []Syntax {
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

func (d *DeclExpr) String() string {
	if d == nil {
		return "nil DeclExpr"
	}
	if d.Init != nil {
		return fmt.Sprintf("DeclExpr<%d>{%s, %s} = %s", d.Id, d.Name.String(), d.Type, d.Init)
	} else {
		return fmt.Sprintf("DeclExpr<%d>{%s, %s}", d.Id, d.Name.String(), d.Type)
	}
}

func (x *DeclExpr) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "DeclExpr"
	}
	return json.Marshal(*x)

}
func (x *DeclExpr) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
