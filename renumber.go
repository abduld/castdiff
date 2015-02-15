package main

import (
	. "github.com/abduld/castdiff/cc"
)

func renumber(prog *Prog) {
	var id int = 0
	nextId := func() int {
		id++
		return id
	}
	Preorder(prog, func(x Syntax) {
		switch x := x.(type) {
		case *Decl:
			x.Id = nextId()
		case *Init:
			x.Id = nextId()
		case *Type:
			x.Id = nextId()
		case *Expr:
			x.Id = nextId()
		case *Stmt:
			x.Id = nextId()
		case *Label:
			x.Id = nextId()
		}
	})
}
