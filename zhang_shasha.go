package main

import (
	"fmt"
	. "github.com/abduld/castdiff/cc"
	_ "runtime/debug"
)

func findKeyRoots(prog *Prog) map[int]Expr {
	Preorder(prog, func(x Syntax) {
		switch x := x.(type) {
		case *Decl:
			if x != nil {
				out := x.String()
				fmt.Println("got decl ", out)
			}
		}
	})
	return nil
}

func ASTDistance(p1, p2 *Prog) int {
	findKeyRoots(p1)
	return 0
}
