package graph_algo

import (
	"fmt"
	"sort"
	"testing"
)

func maxStability(n int, edges [][]int, k int) int {

	type Vertex struct {
		id   int
		rank int
		p    *Vertex
	}

	type Edge struct {
		id     int
		u      int
		v      int
		weight int
	}

	link := func(x, y *Vertex) {
		if x.rank > y.rank {
			y.p = x
		} else {
			x.p = y
			if x.rank == y.rank {
				y.rank = y.rank + 1
			}
		}
	}

	var findSet func(v *Vertex) *Vertex
	findSet = func(v *Vertex) *Vertex {
		if v != v.p {
			v.p = findSet(v.p)
		}
		return v.p
	}

	isDisjoint := func(u, v *Vertex) bool {
		return findSet(u) != findSet(v)
	}

	union := func(x, y *Vertex) {
		px := findSet(x)
		py := findSet(y)
		link(px, py)
	}

	mustList := make([]*Edge, 0, 0)
	optionalList := make([]*Edge, 0, 0)

	for idx := 0; idx < len(edges); idx++ {
		e := &Edge{
			id:     idx,
			u:      edges[idx][0],
			v:      edges[idx][1],
			weight: edges[idx][2],
		}
		if edges[idx][3] == 1 {
			mustList = append(mustList, e)
		} else {
			optionalList = append(optionalList, e)
		}
	}

	vertexList := make([]*Vertex, n, n)

	for idx := 0; idx < n; idx++ {
		v := &Vertex{id: idx, rank: 0, p: nil}
		v.p = v
		vertexList[idx] = v
	}

	mustMin := 2147483647
	// include the must list
	for _, edge := range mustList {
		uVertex := vertexList[edge.u]
		vVertex := vertexList[edge.v]
		if isDisjoint(uVertex, vVertex) {

			if mustMin > edge.weight {
				mustMin = edge.weight
			}
			union(vertexList[edge.u], vertexList[edge.v])

		} else { // connected cycle
			//fmt.Println("# found a cycle for must-include condition")
			return -1
		}
	}

	// we are after max-spanning tree not min-spanning tree
	sort.Slice(optionalList,
		func(idx, jdx int) bool {
			return optionalList[idx].weight > optionalList[jdx].weight
		})

	acceptedList := make([]*Edge, 0, 0)
	for _, edge := range optionalList {
		uVertex := vertexList[edge.u]
		vVertex := vertexList[edge.v]
		if isDisjoint(uVertex, vVertex) {
			union(uVertex, vVertex)
			acceptedList = append(acceptedList, edge)
		}
	}

	for _, v := range vertexList {
		if findSet(vertexList[0]) != findSet(v) {
			//fmt.Println("# not connected vertex found", vertexList[0], v)
			return -1
		}
	}

	l := len(acceptedList)
	if l == 0 {
		return mustMin
	}

	optionalMin := 2147483647
	if k == 0 {
		optionalMin = acceptedList[0].weight
	}

	kdx := 0
	for idx := l - 1; idx >= 0; idx-- {
		w := acceptedList[idx].weight
		if kdx < k {
			w = w * 2
		}
		if optionalMin > w {
			optionalMin = w
		}
		kdx++
	}

	//fmt.Println("mst-min", mustMin, "opt-max", optionalMin)
	if mustMin < optionalMin {
		return mustMin
	} else {
		return optionalMin
	}
}

func TestStabilityMst(t *testing.T) {

	n := 5
	edges := [][]int{
		{0, 1, 24819, 0},
		{2, 4, 86210, 1},
		{1, 2, 53407, 0},
		{3, 4, 56877, 0},
		{1, 3, 89383, 0},
	}
	k := 4
	result := maxStability(n, edges, k)
	fmt.Println("result", result)
	/* n := 4
	edges := [][]int{
		{0, 1, 1, 1},
		{1, 2, 1, 1},
		{2, 3, 1, 1},
		{0, 3, 1, 0},
	}
	k := 2
	result := maxStability(n, edges, k)
	fmt.Println(result) */
}

/*
for idx := 0; idx < k; idx++ {
		optionalList[idx].weight = optionalList[idx].weight * 2
	}

	sort.Slice(optionalList,
		func(idx, jdx int) bool {
			return optionalList[idx].weight < optionalList[jdx].weight
		})
*/
