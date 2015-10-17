package ast

import "encoding/json"

type LabeledStmt struct {
	SyntaxInfo
	Kind  string `json:"kind"`
	Id    int    `json:"id"`
	Labels []*Label  `json:"labels"`
	Expr  Expr   `json:"expr"`
}

func (x *LabeledStmt) GetId() int {
	return x.Id
}

func (x *LabeledStmt) String() string {
	str := ""
	for _, lbl := range x.Labels {
		str += lbl.String() + "\n"
	}
	return str + x.Expr.String()
}

func (x *LabeledStmt) GetChildren() []Syntax {
	children := make([]Syntax, len(x.Labels) + 1)
	for ii, lbl := range x.Labels {
		children[ii] = lbl
	}
	children[len(x.Labels)] = x.Expr
	return children
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

func (x *LabeledStmt) IsStmt() bool {
	return true
}