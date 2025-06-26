package graph_algo

import (
	"fmt"
	"testing"
)

func findCircleNum(edges [][]int) int {
	n := len(edges)

	type vertex struct {
		val  int
		rank int
		p    *vertex
	}

	// makeset
	vertexList := make([]*vertex, n, n)
	for idx := 0; idx < n; idx++ {
		v := &vertex{idx + 1, 0, nil}
		v.p = v
		vertexList[idx] = v
	}

	// link ops
	var link func(x, y *vertex)
	link = func(x, y *vertex) {
		if x.rank > y.rank {
			y.p = x
		} else {
			x.p = y
			if x.rank == y.rank {
				y.rank = y.rank + 1
			}
		}
	}

	var findSet func(v *vertex) *vertex
	findSet = func(v *vertex) *vertex {
		if v != v.p {
			v.p = findSet(v.p)
		}
		return v.p
	}

	//union
	var union func(x, y *vertex)
	union = func(x, y *vertex) {
		link(findSet(x), findSet(y))
	}

	for idx := 0; idx < n; idx++ {
		for jdx := idx + 1; jdx < n; jdx++ {
			if edges[idx][jdx] == 1 {
				union(vertexList[idx], vertexList[jdx])

			}
		}
	}

	setMap := make(map[int]bool)

	for idx := range vertexList {
		v := vertexList[idx]
		p := findSet(v)
		setMap[p.val] = true
	}

	return len(setMap)

}

func TestFindCircleNum(t *testing.T) {
	edges := [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}
	ans := findCircleNum(edges)
	fmt.Println(ans)
}
