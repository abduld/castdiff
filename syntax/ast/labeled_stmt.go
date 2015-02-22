package ast

import "encoding/json"

type LabeledStmt struct {
	SyntaxInfo
	Kind  string `json:"kind"`
	Id    int    `json:"id"`
	Label Label  `json:"label"`
	Expr  Expr   `json:"expr"`
}

func (x *LabeledStmt) GetId() int {
	return x.Id
}

func (x *LabeledStmt) String() string {
	return x.Label.String() + ": " + x.Expr.String()
}

func (x *LabeledStmt) GetChildren() []Syntax {
	return []Syntax{x.Label, x.Expr}
}

func (x *LabeledStmt) ToSymbolLiteral() *SymbolLiteral {
	return x
}

func (x *LabeledStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "LabeledStmt"
	}
	return json.Marshal(*x)

}
func (x *LabeledStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
