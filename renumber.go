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
		case nil:
			return
		case *EmptyLiteral:
			x.Id = nextId()
		case *BooleanLiteral:
			x.Id = nextId()
		case *IntegerLiteral:
			x.Id = nextId()
		case *CharLiteral:
			x.Id = nextId()
		case *RealLiteral:
			x.Id = nextId()
		case *StringLiteral:
			x.Id = nextId()
		case *SymbolLiteral:
			x.Id = nextId()
		case *LanguageKeyword:
			x.Id = nextId()
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
