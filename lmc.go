package main

import (
	_ "fmt"

	. "github.com/abduld/castdiff/cc"
)

func leftMostChild(prog *Prog) map[int]Syntax {
	lmc := map[int]Syntax{}
	Postorder(prog, func(x Syntax) {
		switch x := x.(type) {
		case *EmptyLiteral:
			lmc[x.Id] = x
		case *BooleanLiteral:
			lmc[x.Id] = x
		case *IntegerLiteral:
			lmc[x.Id] = x
		case *CharLiteral:
			lmc[x.Id] = x
		case *RealLiteral:
			lmc[x.Id] = x
		case *StringLiteral:
			lmc[x.Id] = x
		case *SymbolLiteral:
			lmc[x.Id] = x
		case *LanguageKeyword:
			lmc[x.Id] = x
		case *Decl:
			if x.Init != nil {
				lmc[x.Id] = lmc[x.Init.Id]
			} else if x.Type.Kind == Func {
				lmc[x.Id] = x
			} else if x.Body != nil {
				lmc[x.Id] = lmc[x.Body.Id]
			} else if x.Name.String() == "" {
				lmc[x.Id] = lmc[x.Type.Id]
			} else {
				lmc[x.Id] = x
			}
		case *Init:
			lmc[x.Id] = lmc[x.Expr.Id]
		case *Expr:
			switch x.Op {
			case Cond:
				fallthrough
			case Comma:
				lmc[x.Id] = lmc[x.List[0].Id]
			default:
				if x.Left != nil {
					lmc[x.Id] = lmc[x.Left.Id]
				} else {
					lmc[x.Id] = x
				}
			}
		case *Type:
			switch x.Kind {
			case Func:
				lmc[x.Id] = lmc[x.Base.Id]
			default:
				if len(x.Decls) == 0 {
					lmc[x.Id] = x
				} else {
					lmc[x.Id] = lmc[x.Decls[0].Id]
				}
			}
		case *Stmt:
			switch x.Op {
			case Block:
				blk := x.Block
				if len(blk) == 0 {
					lmc[x.Id] = x
				} else {
					lmc[x.Id] = lmc[blk[0].Id]
				}
			case Break:
				lmc[x.Id] = x
			case Continue:
				lmc[x.Id] = x
			case Do:
				lmc[x.Id] = lmc[x.Body.Id]
			case Empty:
				lmc[x.Id] = x
			case For:
				lmc[x.Id] = lmc[x.Pre.Id]
			case If:
				lmc[x.Id] = lmc[x.Expr.Id]
			case Goto:
				lmc[x.Id] = x
			case Return:
				if x.Expr == nil {
					lmc[x.Id] = x
				} else {
					lmc[x.Id] = lmc[x.Expr.Id]
				}
			case StmtDecl:
				lmc[x.Id] = lmc[x.Decl.Id]
			case StmtExpr:
				lmc[x.Id] = lmc[x.Expr.Id]
			case Switch:
				lmc[x.Id] = lmc[x.Expr.Id]
			case While:
				lmc[x.Id] = lmc[x.Expr.Id]
			}
		}
	})
	/*
		for k, v := range lmc {
			fmt.Println(k, "  : ", v)
		}
		fmt.Println("-------------")
	*/
	return lmc
}

func leftMostChild0(prog *Prog) map[int]Syntax {
	lmc := map[int]Syntax{}
	Postorder(prog, func(x Syntax) {
		children := x.GetChildren()
		if len(children) == 0 {
			lmc[x.GetId()] = x
		} else {
			lmc[x.GetId()] = lmc[children[0].GetId()]
		}
	})
	return lmc
}
