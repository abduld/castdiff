package ast

import "encoding/json"

type Label struct {
	SyntaxInfo
	Kind string         `json:"kind"`
	Id   int            `json:"id"`
	IsCase bool `json:"is_case`
	Name Expr `json:"name"`
}

func (x *Label) GetId() int {
	return x.Id
}

func (x *Label) String() string {
	if x.IsCase {
		return "case " + x.Name.String() + ":"
	}
	return x.Name.String() + ":"
}

func (x *Label) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *Label) MarshalJSON() ([]byte, error) {
	if x != nil {
		x.Kind = "Label"
	}
	return json.Marshal(*x)

}
func (x *Label) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, x)
}
