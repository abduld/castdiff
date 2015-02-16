package main

import (
	. "github.com/abduld/castdiff/cc"
)

func leafCount(prog *Prog) int {
	var n int = 0
	Postorder(prog, func(x Syntax) {
		children := x.GetChildren()
		if len(children) == 0 {
			n++
		}
	})
	return n
}
