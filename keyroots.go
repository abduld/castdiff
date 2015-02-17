package main

import (
	"sort"

	. "github.com/abduld/castdiff/cc"
)

/* A key root is a node of T that either has a left
 * sibling or is the root of T
 */
func keyroots(prog *Prog) []Syntax {
	lmc := leftMostChilden(prog)
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
	return skr
}
