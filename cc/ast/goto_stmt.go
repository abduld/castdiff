package ast

import "encoding/json"

type GotoStmt struct {
	SyntaxInfo
	Kind   string `json:"kind"`
	Id     int    `json:"id"`
	Target *Label `json:"cond"`
}

func (x *GotoStmt) GetId() int {
	return x.Id
}

func (x *GotoStmt) String() string {
	return "goto " + x.Target.String()
}

func (x *GotoStmt) GetChildren() []Syntax {
	return []Syntax{x.Target}
}

func (x *GotoStmt) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "GotoStmt"
	}
	return json.Marshal(*x)

}
func (x *GotoStmt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
