package ast

import "fmt"

// Walk traverses the syntax x, calling before and after on entry to and exit from
// each Syntax encountered during the traversal. In case of cross-linked input,
// the traversal never visits a given Syntax more than once.
func Walk(x Syntax, before, after func(Syntax)) {
	seen := map[Syntax]bool{
		nil:                     true,
		(*EmptyLiteral)(nil):    true,
		(*IntegerLiteral)(nil):  true,
		(*CharLiteral)(nil):     true,
		(*RealLiteral)(nil):     true,
		(*StringLiteral)(nil):   true,
		(*SymbolLiteral)(nil):   true,
		(*LanguageKeyword)(nil): true,
		(*DeclStmt)(nil):        true,
		(*Init)(nil):            true,
		(*Type)(nil):            true,
		(Expr)(nil):             true,
		(Stmt)(nil):             true,
		(*Label)(nil):           true,
	}
	walk(x, before, after, seen)
}

func walk(x Syntax, before, after func(Syntax), seen map[Syntax]bool) {
	if x == nil {
		return
	}
	if seen[x] {
		return
	}
	seen[x] = true
	before(x)
	switch x := x.(type) {
	default:
		panic(fmt.Sprintf("walk: unexpected type %T", x))

	case *EmptyLiteral:
		//ok
	case *BooleanLiteral:
		//ok
	case *IntegerLiteral:
		//ok
	case *CharLiteral:
		//ok
	case *RealLiteral:
		//ok
	case *StringLiteral:
		//ok
	case *SymbolLiteral:
		//ok
	case *LanguageKeyword:
		//ok
	/*
		case *Prog:
			for _, d := range x.Decls {
				walk(d, before, after, seen)
			}
	*/
	case *DeclStmt:
		walk(x.Type, before, after, seen)
		walk(x.Name, before, after, seen)
		walk(x.Init, before, after, seen)

	case *Init:
		for _, b := range x.Braced {
			walk(b, before, after, seen)
		}
		walk(x.Expr, before, after, seen)

	case *Type:
		walk(x.Base, before, after, seen)
		for _, d := range x.Decls {
			walk(d, before, after, seen)
		}
		//walk(x.Width, before, after, seen)
		/*
					case *Expr:
						walk(x.Left, before, after, seen)
						walk(x.Text, before, after, seen)
						walk(x.Right, before, after, seen)
						for _, y := range x.LaunchParams {
							walk(y, before, after, seen)
						}
						for _, y := range x.Texts {
							walk(y, before, after, seen)
						}
						for _, y := range x.List {
							walk(y, before, after, seen)
						}
						walk(x.Type, before, after, seen)
						walk(x.Init, before, after, seen)
						for _, y := range x.Block {
							walk(y, before, after, seen)
						}
			case *Stmt:
				walk(x.Pre, before, after, seen)
				walk(x.Expr, before, after, seen)
				walk(x.Post, before, after, seen)
				walk(x.Decl, before, after, seen)
				walk(x.Body, before, after, seen)
				walk(x.Else, before, after, seen)
				walk(x.Text, before, after, seen)
				for _, y := range x.Block {
					walk(y, before, after, seen)
				}
				for _, y := range x.Labels {
					walk(y, before, after, seen)
				}

		*/
	case *Label:
		walk(x.Name, before, after, seen)
	}
	after(x)
}

// Preorder calls f for each piece of syntax of x in a preorder traversal.
func Preorder(x Syntax, f func(Syntax)) {
	Walk(x, f, func(Syntax) {})
}

// Preorder calls f for each piece of syntax of x in a postorder traversal.
func Postorder(x Syntax, f func(Syntax)) {
	Walk(x, func(Syntax) {}, f)
}
