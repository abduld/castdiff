package ast

import "encoding/json"

type FuncStmt struct {
	SyntaxInfo
	Kind       string         `json:"kind"`
	Id         int            `json:"id"`
	ReturnType *Type          `json:"rettype"`
	Name       *SymbolLiteral `json:"name"`
	Args       []DeclStmt     `json:"args"`
	Body       Syntax         `json:"body"`
	IsDecl     bool           `json:"decl"`
}

func (x *FuncStmt) GetId() int {
	return x.Id
}

func (x *FuncStmt) String() string {
	ret := ""
	ret += x.RetType.String() + " " + x.Name.String() + "( "
	for _, elem := range x.Args {
		ret += elem.String() + ", "
	}
	ret += ")"
	if x.IsDecl {
		ret += "{\n" + x.Body.String() + "}"
	} else {
		ret += ";"
	}
	return ret
}

func (x *FuncStmt) GetChildren() []Syntax {
	children := []Syntax{x.RetType, x.Name}
	for _, arg := range x.Args {
		children = append(children, &arg)
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
