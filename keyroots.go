package main

import (
	"fmt"
	. "github.com/abduld/castdiff/cc"
	"sort"
)

func numLeaves(lmc map[int]Syntax) int {
	var n int = 0
	for k, v := range lmc {
		if k == v.GetId() {
			n++
		}
	}
	return n
}

/* A key root is a node of T that either has a left
 * sibling or is the root of T
 */
func keyroots(prog *Prog) []Syntax {
	lmc := leftMostChild(prog)
	k := numLeaves(lmc)
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
	for _, v := range skr {
		fmt.Println(v)
	}
	return skr
}
