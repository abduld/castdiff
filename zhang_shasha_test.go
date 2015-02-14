package main

import (
	cc "github.com/abduld/castdiff/cc"
	"testing"
)

func TestDistanceDecl(t *testing.T) {
	p1, _ := cc.ParseProg("int x;")
	p2, _ := cc.ParseProg("int x;")
	if ASTDistance(p1, p2) != 0 {
		t.Errorf("Distance between %#q, %#q was not expected", p1, p2)
	}
}
