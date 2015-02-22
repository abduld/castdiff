package ast

import "encoding/json"

type Label struct {
	SyntaxInfo
	Kind  string       `json:"kind"`
	Id    int          `json:"id"`
	Name SymbolLiteral `json:"label"`
}

func (x *Label) GetId() int {
	return x.Id
}

func (x *Label) String() string {
	return x.Name.String()
}

func (x *Label) GetChildren() []Syntax {
	return []Syntax{}
}

func (x *Label) ToSymbolLiteral() *SymbolLiteral {
	return &x.Name
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
