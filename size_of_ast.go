package main

import (
	. "github.com/abduld/castdiff/cc"
)

func sizeOfAST(prog *Prog) int {
	n := 0
	Postorder(prog, func(x Syntax) {
		n++
	})
	return n
}
