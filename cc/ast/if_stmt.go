package ast

import "encoding/json"

type IfStmt struct {
	SyntaxInfo
	Kind string `json:"kind"`
	Id   int    `json:"id"`
	Cond Expr   `json:"cond"`
	Then Expr   `json:"then"`
	Else Expr   `json:"else"`
}

func (x *IfStmt) GetId() int {
	return x.Id
}

func (x *IfStmt) String() string {
	if x.Else == nil {
		return "if (" + x.Cond.String() + ") \n\t" + x.Then.String()
	} else {
		return "if (" + x.Cond.String() + ") \n\t" + x.Then.String() + "\n else \n\t" + x.Else.String()
	}
}

func (x *IfStmt) GetChildren() []Syntax {
	if x.Else == nil {
		return []Syntax{x.Cond, x.Then}
	} else {
		return []Syntax{x.Cond, x.Then, x.Else}
	}
}

func (x *IfStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "IfStmt"
	}
	return json.Marshal(*x)

}
func (x *IfStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}

func (x *IfStmt) IsBlock() bool {
	return true
}

func (x *IfStmt) IsStmt() bool {
	return true
}