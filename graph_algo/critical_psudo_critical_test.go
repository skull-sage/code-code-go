package graph_algo

import (
	"fmt"
	"sort"
	"testing"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {

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

	edgeList := make([]*Edge, len(edges), len(edges))
	for idx := 0; idx < len(edges); idx++ {
		edgeList[idx] = &Edge{
			id:     idx,
			u:      edges[idx][0],
			v:      edges[idx][1],
			weight: edges[idx][2],
		}
	}

	sort.Slice(edgeList, func(idx, jdx int) bool {
		return edgeList[idx].weight < edgeList[jdx].weight
	})
	vertexList := make([]*Vertex, n, n)

	kruskal := func(edgeList []*Edge, keepEdge *Edge, ignoreEdge *Edge) int {

		// MAKE_SET of disjoint set
		for idx := 0; idx < n; idx++ {
			v := &Vertex{id: idx, rank: 0, p: nil}
			v.p = v
			vertexList[idx] = v
		}
		totalWeight := 0

		if keepEdge != nil {
			totalWeight += keepEdge.weight
			union(vertexList[keepEdge.u], vertexList[keepEdge.v])

		}

		for _, edge := range edgeList {
			if edge == ignoreEdge {
				continue
			}
			uVertex := vertexList[edge.u]
			vVertex := vertexList[edge.v]

			if isDisjoint(uVertex, vVertex) {
				totalWeight += edge.weight
				union(uVertex, vVertex)

			}

		}
		return totalWeight
	}

	minCost := kruskal(edgeList, nil, nil)

	criticalList := make([]int, 0, 0)
	pseudoCriticalList := make([]int, 0, 0)

	var cost int
	for _, edge := range edgeList {

		cost = kruskal(edgeList, nil, edge)
		uVert := vertexList[edge.u]
		vVert := vertexList[edge.v]
		if isDisjoint(uVert, vVert) || cost > minCost {
			criticalList = append(criticalList, edge.id)
		} else {
			cost = kruskal(edgeList, edge, nil)
			if cost == minCost {
				pseudoCriticalList = append(pseudoCriticalList, edge.id)
			}
		}

	}

	return [][]int{criticalList, pseudoCriticalList}
}

func TestCritical(t *testing.T) {
	/* n := 5
	edges := [][]int{
		{0, 1, 1},
		{1, 2, 1},
		{2, 3, 2},
		{0, 3, 2},
		{0, 4, 3},
		{3, 4, 3},
		{1, 4, 6},
	} */

	 n := 3
	edges := [][]int{
		{0, 1, 1},
		{0, 2, 2},
		{1, 2, 3},
	} 

	result := findCriticalAndPseudoCriticalEdges(n, edges)
	fmt.Println(result)
}
