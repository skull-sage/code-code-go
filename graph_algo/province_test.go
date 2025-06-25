package graph_algo

import "testing"

func findCircleNum(edges [][]int) int {
	n := len(edges)

	type vertex struct {
		val  int
		rank int
		p    int
	}

	// makeset
	vertexList := make([]vertex, n, n)
	for idx := 0; idx < n; idx++ {
		v := vertexList[idx]
		v.val = idx
		v.rank = 0
		v.p = v.val
	}

	var link func(x, y vertex)
	link = func(x, y vertex) {
		if x.rank > y.rank {
			y.p = x.val
		} else {
			x.p = y.val
			if x.rank == y.rank {
				y.rank = y.rank + 1
			}
		}
	}

	for idx := 0; idx < n; idx++ {
		for jdx := idx + 1; jdx < n; jdx++ {
			if edges[idx][jdx] == 1 {

			}
		}
	}
}

func TestFindCircleNum(t *testing.T) {
	edges := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	ans := findCircleNum(edges)
}
