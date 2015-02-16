package main

import (
	"fmt"
	. "github.com/abduld/castdiff/cc"
	"sort"
)

/* A key root is a node of T that either has a left
 * sibling or is the root of T
 */
func keyroots(prog *Prog) []Syntax {
	lmc := leftMostChild(prog)
	k := leafCount(prog)
	visitedQ := map[int]bool{}
	kr := make([]int, k)
	ii := len(lmc)
	for k > 0 {
		if _, ok := visitedQ[lmc[ii].GetId()]; !ok {
			k--
			kr[k] = lmc[ii].GetId()
			visitedQ[lmc[ii].GetId()] = true
		}
		ii--
	}
	sort.Ints(kr)
	skr := make([]Syntax, len(kr))
	for i, v := range kr {
		skr[i] = lmc[v]
	}
	fmt.Println("leaf count = ", leafCount(prog))
	for k, v := range lmc {
		fmt.Println(k, " >= ", v)
	}
	for k, v := range skr {
		fmt.Println(k, " = ", v)
	}
	fmt.Println(prog.Decls[0])
	return skr
}
