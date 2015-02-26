package ast

import (
	"encoding/json"
	"strings"
)

type BlockStmt struct {
	SyntaxInfo
	Kind  string `json:"kind"`
	Id    int    `json:"id"`
	Stmts []Stmt `json:"stmts"`
}

func (x *BlockStmt) GetId() int {
	return x.Id
}

func (x *BlockStmt) String() string {
	sstmts := make([]string, len(x.Stmts))
	for ii, stmt := range x.Stmts {
		sstmts[ii] = stmt.String()
	}
	return strings.Join(sstmts, ";\n")
}

func (x *BlockStmt) GetChildren() []Syntax {
	children := make([]Syntax, len(x.Stmts))
	for ii, stmt := range x.Stmts {
		children[ii] = stmt
	}
	return children
}

func (x *BlockStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "BlockStmt"
	}
	return json.Marshal(*x)

}
func (x *BlockStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
func (b *BlockStmt) IsStmt() bool {
	return true
}
