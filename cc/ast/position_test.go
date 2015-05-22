package ast

import (
	"encoding/json"
	"testing"
)

func TestPosition(t *testing.T) {
	p := &Position{File: "test.c", Line: 1, Byte: 1}
	js, _ := json.Marshal(&p)
	t.Log(string(js))
	if p.Line != 1 {
		t.Error()
	}
}
