package ast

import "encoding/json"

type ForStmt struct {
	SyntaxInfo
	Kind   string `json:"kind"`
	Id     int    `json:"id"`
	Init   Expr   `json:"init"`
	Cond   Expr   `json:"cond"`
	Update Expr   `json:"update"`
	Body   Expr   `json:"body"`
}

func (x *ForStmt) GetId() int {
	return x.Id
}

func (x *ForStmt) String() string {
	return "for ( " + x.Init.String() + "; " +
		x.Cond.String() + "; " +
		x.Update.String() + "; " + " ) {\n" +
		x.Body.String() + "}"
}

func (x *ForStmt) GetChildren() []Syntax {
	return []Syntax{x.Init, x.Cond, x.Update, x.Body}
}

func (x *ForStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "ForStmt"
	}
	return json.Marshal(*x)

}
func (x *ForStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}

func (x *ForStmt) IsStmt() bool {
	return true
}