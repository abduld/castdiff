package ast

import "encoding/json"

type SwitchStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Cond Expr   `json:"cond"`
	Body Expr   `json:"body"`
}

func (x *SwitchStmt) GetId() int {
	return x.Id
}

func (x *SwitchStmt) String() string {
	return "switch ( " + x.Cond.String() + ") {\n\t" + x.Body.String() + "\n}"
}

func (x *SwitchStmt) GetChildren() []Syntax {
	return []Syntax{x.Cond, x.Body}
}

func (x *SwitchStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "SwitchStmt"
	}
	return json.Marshal(*x)

}
func (x *SwitchStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}

func (x *SwitchStmt) IsBlock() bool {
	return true
}

func (x *SwitchStmt) IsStmt() bool {
	return true
}