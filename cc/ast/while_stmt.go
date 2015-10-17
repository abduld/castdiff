package ast

import "encoding/json"

type WhileStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Cond Expr   `json:"cond"`
	Body Expr   `json:"body"`
}

func (x *WhileStmt) GetId() int {
	return x.Id
}

func (x *WhileStmt) String() string {
	return "while ( " + x.Cond.String() + ") {\n\t" + x.Body.String() + "\n}"
}

func (x *WhileStmt) GetChildren() []Syntax {
	return []Syntax{x.Cond, x.Body}
}

func (x *WhileStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "WhileStmt"
	}
	return json.Marshal(*x)

}
func (x *WhileStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}

func (x *WhileStmt) IsBlock() bool {
	return true
}

func (x *WhileStmt) IsStmt() bool {
	return true
}