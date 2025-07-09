package graph_algo

import (
	"fmt"
	"testing"
)

func findRedundantConnection(edges [][]int) []int {
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

	/* //union
	var union func(x, y *vertex)
	union = func(x, y *vertex) {
		link(findSet(x), findSet(y))
	} */

	exclude := make([][]int, 0, 0)
	for idx := 0; idx < n; idx++ {
		edge := edges[idx]
		x, y := edge[0], edge[1]
		setX := findSet(vertexList[x-1])
		setY := findSet(vertexList[y-1])
		// setX, setY is essentially compressed-path parent of x & y
		if setX == setY {
			exclude = append(exclude, edge)
		} else {
			link(setX, setY)
		}

	}

	return exclude[len(exclude)-1]
}

func TestRedundant(t *testing.T) {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	result := findRedundantConnection(edges)
	fmt.Println(result)

}
