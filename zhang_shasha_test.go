package main

import (
	_ "fmt"
	cc "github.com/abduld/castdiff/cc"
	"testing"
)

func TestDistanceDecl(t *testing.T) {
	p1, err := cc.ParseProg("int ii, jj; void f() { return 0; }")
	if err != nil {
		t.Errorf("Unable to parse p1 -- %#q", err)
		return
	}
	/*
		p2, err := cc.ParseProg("int x;")
		if err != nil {
			t.Errorf("Unable to parse p2 -- %#q", err)
			return
		}
	*/
	if ASTDistance(p1, p1) != 0 {
		t.Errorf("Distance between %#q, %#q was not expected", p1, p1)
	}
}
