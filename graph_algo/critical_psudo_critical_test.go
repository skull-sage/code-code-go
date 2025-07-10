package graph_algo

import (
	"fmt"
	"sort"
	"testing"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {

	type vertex struct {
		id   int
		rank int
		p    *vertex
	}

	type edge struct {
		id     int
		u      *vertex
		v      *vertex
		weight int
	}

	link := func(x, y *vertex) {
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

	isDisjoint := func(u, v *vertex) bool {
		return findSet(u) != findSet(v)
	}

	union := func(x, y *vertex) {
		px := findSet(x)
		py := findSet(y)
		link(px, py)
	}

	vertexList := make([]*vertex, n, n)
	for idx := 0; idx < n; idx++ {
		v := &vertex{id: idx, rank: 0, p: nil}
		v.p = v
		vertexList[idx] = v
	}

	edgeList := make([]*edge, len(edges), len(edges))
	for idx := 0; idx < len(edges); idx++ {
		edgeList[idx] = &edge{
			id:     idx,
			u:      vertexList[edges[idx][0]],
			v:      vertexList[edges[idx][1]],
			weight: edges[idx][2],
		}
	}

	sort.Slice(edgeList, func(idx, jdx int) bool {
		return edgeList[idx].weight < edgeList[jdx].weight
	})

	kruskal := func(edgeList []*edge, ignoreEdge *edge) int {

		totalWeight := 0
		for _, edge := range edgeList {
			if edge == ignoreEdge {
				continue
			}
			if isDisjoint(edge.u, edge.v) {
				totalWeight += edge.weight
				union(edge.u, edge.v)
			}
		}
		return totalWeight
	}

	minCost := kruskal(edgeList, nil)
	fmt.Printf("# minCost: %d\n", minCost)

	criticalList := make([]int, 0, 0)
	pseudoCriticalList := make([]int, 0, 0)

	for _, edge := range edgeList {
		cost := kruskal(edgeList, edge)
		fmt.Printf("# ignoring edge: %d cost=%d\n", edge.id, cost)
		if isDisjoint(edge.u, edge.v) || cost > minCost {
			criticalList = append(criticalList, edge.id)
		} else if cost == minCost {
			pseudoCriticalList = append(pseudoCriticalList, edge.id)
		}
	}

	return [][]int{criticalList, pseudoCriticalList}
}

func TestCritical(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1, 1},
		{1, 2, 1},
		{2, 3, 2},
		{0, 3, 2},
		{0, 4, 3},
		{3, 4, 3},
		{1, 4, 6},
	}

	result := findCriticalAndPseudoCriticalEdges(n, edges)
	fmt.Println(result)
}
