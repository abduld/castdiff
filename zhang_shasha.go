package main

import (
	"fmt"
	_ "runtime/debug"

	. "github.com/abduld/castdiff/cc"
)

func populateHashMap(prog *Prog) map[int]Syntax {
	hm := map[int]Syntax{}
	Preorder(prog, func(x Syntax) {
		switch x := x.(type) {
		case *Decl:
			hm[x.Id] = x
		case *Init:
			hm[x.Id] = x
		case *Type:
			hm[x.Id] = x
		case *Expr:
			hm[x.Id] = x
		case *Stmt:
			hm[x.Id] = x
		case *Label:
			hm[x.Id] = x
		}
	})
	return hm
}

func ASTDistance(p1, p2 *Prog) int {
	renumber(p1)
	keyroots(p1)
	//hm := populateHashMap(p1)
	fmt.Println("done")
	return 0
}
