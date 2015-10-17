package ast

import "encoding/json"

type DoStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Cond Expr   `json:"cond"`
	Body Expr   `json:"body"`
}

func (x *DoStmt) GetId() int {
	return x.Id
}

func (x *DoStmt) String() string {
	return "do {\n\t" + x.Body.String() + "\n} while { " + x.Cond.String()
}

func (x *DoStmt) GetChildren() []Syntax {
	return []Syntax{x.Cond, x.Body}
}

func (x *DoStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "DoStmt"
	}
	return json.Marshal(*x)

}
func (x *DoStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
