package ast

import (
	"encoding/json"
)

type Position struct {
	Kind string `json:"kind"`
	File string `json:"file"`
	Line int    `json:"line"`
	Byte int    `json:"byte"`
}

func (p *Position) MarshalJSON() ([]byte, error) {
	if p != nil {
		p.Kind = "Position"
	}
	return json.Marshal(*p)
}

func (p *Position) UnmarshalJSON(text []byte) error {
	return json.Unmarshal(text, p)
}
