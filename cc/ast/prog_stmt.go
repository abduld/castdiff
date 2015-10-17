package ast

import (
	"encoding/json"
	"strings"
)

type Prog struct {
	SyntaxInfo
	Kind  string `json:"kind"`
	Id    int
	Decls []Stmt
}

func (x *Prog) GetId() int {
	return x.Id
}

func (x *Prog) GetChildren() []Syntax {
	lst := []Syntax{}
	for _, elem := range x.Decls {
		if elem != nil {
			lst = append(lst, elem)
		}
	}
	return lst
}

func (x *Prog) String() string {
	children := x.GetChildren()
	sstmts := make([]string, len(children))
	for ii, child := range children {
		sstmts[ii] = child.String()
	}
	return strings.Join(sstmts, ";\n")

}

func (x *Prog) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "Program"
	}
	return json.Marshal(*x)

}
func (x *Prog) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
func (x *Prog) IsStmt() bool {
	return true
}

func (x *Prog) IsBlock() bool {
	return true
}
