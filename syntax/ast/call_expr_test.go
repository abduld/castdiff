package ast

import (
	"encoding/json"
	"testing"
)

func TestCallExpr(t *testing.T) {
	p := &CallExpr{
		Callee: &SymbolLiteral{Value: "test"},
		Args: []Expr{
			&SymbolLiteral{Value: "arg1"},
			&SymbolLiteral{Value: "arg2"},
		},
	}
	js, _ := json.Marshal(&p)
	t.Log(string(js))
}
