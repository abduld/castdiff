package main

import (
	"fmt"
	_ "runtime/debug"

	. "github.com/abduld/castdiff/cc"
)

func imax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func imin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func fmin(a, b float32) float32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func fmin3(a, b, c float32) float32 {
	return fmin(a, fmin(b, c))
}

func forestDistance(
	deletePenatly, insertPenatly, renamePenalty float32,
	kr1, kr2 []Syntax,
	x, y int,
	l1, l2 map[int]Syntax,
	td [][]float32,
) {
	i := kr1[x]
	j := kr2[y]
	iid := i.GetId()
	jid := j.GetId()
	yDim := imax(iid-l1[iid].GetId()+1, 1)
	xDim := imax(jid-l2[jid].GetId()+1, 1)
	fd := make([][]float32, yDim)
	for ii := range fd {
		fd[ii] = make([]float32, xDim)
	}
	fd[0][0] = 0
	for di := 1; di < yDim; di++ {
		fd[di][0] = fd[di-1][0] + deletePenatly
	}
	for dj := 1; dj < xDim; dj++ {
		fd[0][dj] = fd[0][dj-1] + insertPenatly
	}
	for di := 1; di < yDim; di++ {
		for dj := 1; dj < xDim; dj++ {
			if l1[di].String() == l1[yDim].String() &&
				l2[dj].String() == l2[xDim].String() {
				fd[di][dj] = fmin3(
					fd[di-1][dj]+deletePenatly,
					fd[di][dj-1]+insertPenatly,
					fd[di-1][dj-1]+renamePenalty,
				)
				td[di+iid-1][dj+jid-1] = 0
			} else {
				fd[di][dj] = fmin3(
					fd[di-1][dj]+deletePenatly,
					fd[di][dj-1]+insertPenatly,
					fd[l1[di].GetId()-l1[iid].GetId()-1][l2[dj].GetId()-l2[jid].GetId()-1]+
						td[di+iid-1][dj+jid-1],
				)
			}
		}
	}
	return
}

func ASTDistance(p1, p2 *Prog) int {
	var deletePenatly, insertPenatly, renamePenalty float32
	deletePenatly = 1
	insertPenatly = 1
	renamePenalty = 1
	renumber(p1)
	renumber(p2)
	t1Size := leafCount(p1)
	t2Size := leafCount(p2)
	td := make([][]float32, t1Size)
	for ii := range td {
		td[ii] = make([]float32, t2Size)
	}
	l1 := leftMostChilden(p1)
	l2 := leftMostChilden(p2)
	kr1 := keyroots(p1)
	kr2 := keyroots(p2)
	for x := range kr1 {
		for y := range kr2 {
			forestDistance(
				deletePenatly, insertPenatly, renamePenalty,
				kr1, kr2,
				x, y,
				l1, l2,
				td,
			)
		}
	}
	fmt.Println(td)
	fmt.Println("edit distance = ", td[t1Size-1][t2Size-1])
	return 0
}
