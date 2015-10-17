package ast

import (
	"encoding/json"
	"testing"
)

func TestIfStmt(t *testing.T) {
	p := &IfStmt{
		Cond: &SymbolLiteral{Value: "condvalue"},
		Then: &SymbolLiteral{Value: "thenvalue"},
		Else: &SymbolLiteral{Value: "elsevalue"},
		}
	js, _ := json.Marshal(&p)
	t.Log(string(js))
}
