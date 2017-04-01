package ast

import (
	"encoding/json"
	"strings"
)

type FuncStmt struct {
	SyntaxInfo
	Kind       string         `json:"kind"`
	Id         int            `json:"id"`
	ReturnType *Type          `json:"rettype"`
	Name       *SymbolLiteral `json:"name"`
	Args       []Stmt         `json:"args"`
	Body       Syntax         `json:"body"`
	Storage    Storage        `json:"storage"`
	IsDecl     bool           `json:"decl"`
}

func (x *FuncStmt) GetId() int {
	return x.Id
}

func (x *FuncStmt) String() string {
	ret := ""
	ret += x.ReturnType.String() + " " + x.Name.String() + "("
	for _, elem := range x.Args {
		ret += elem.String() + ", "
	}
	ret = strings.TrimRight(ret, ", ")
	ret += ")"
	if !x.IsDecl {
		ret += " {\n" + x.Body.String() + "\n}"
	} else {
		ret += ";"
	}
	return ret
}

func (x *FuncStmt) GetChildren() []Syntax {
	children := []Syntax{x.ReturnType, x.Name}
	for _, arg := range x.Args {
		children = append(children, arg)
	}
	return children
}

func (x *FuncStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "FuncStmt"
	}
	return json.Marshal(*x)

}
func (x *FuncStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}

func (x *FuncStmt) IsStmt() bool {
	return true
}
