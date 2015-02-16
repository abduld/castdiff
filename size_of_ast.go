package main

import (
	. "github.com/abduld/castdiff/cc"
)

func sizeOfAST(prog *Prog) int {
	var n int = 0
	Postorder(prog, func(x Syntax) {
		n++
	})
	return n
}
