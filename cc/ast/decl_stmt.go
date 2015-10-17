package ast

import (
	"encoding/json"
	"fmt"
)

type DeclStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id      int            `json:"id"`
	Type    *Type          `json:"type"`
	Storage Storage        `json:"storage"`
	Name    *SymbolLiteral `json:"name"`
	Init    Expr           `json:"init"`
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
func (d *DeclStmt) IsStmt() bool {
	return true
}

func (x *DeclStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "DeclStmt"
	}
	return json.Marshal(*x)

}
func (x *DeclStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
