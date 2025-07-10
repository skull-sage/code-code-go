package graph_algo

import (
	"sort"
)

func minCostConnectPoints(points [][]int) int {

	type vertex struct {
		x    int
		y    int
		rank int
		p    *vertex
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

	n := len(points)
	vertexList := make([]*vertex, n, n)
	type edge struct {
		u      *vertex
		v      *vertex
		weight int
	}

	for idx := 0; idx < n; idx++ {
		v := &vertex{x: points[idx][0], y: points[idx][1], rank: 0, p: nil}
		v.p = v
		vertexList[idx] = v
	}

	edgeList := make([]edge, 0, 0)
	for idx := 0; idx < n; idx++ {
		for jdx := idx + 1; jdx < n; jdx++ {
			u := vertexList[idx]
			v := vertexList[jdx]
			dx := u.x - v.x
			if dx < 0 {
				dx = -dx
			}
			dy := u.y - v.y
			if dy < 0 {
				dy = -dy
			}

			edgeList = append(edgeList, edge{
				u:      u,
				v:      v,
				weight: dx + dy,
			})
		}
	}

	sort.Slice(edgeList, func(idx, jdx int) bool {
		return edgeList[idx].weight <= edgeList[jdx].weight
	})

	cost := 0
	for _, e := range edgeList {
		if isDisjoint(e.u, e.v) {
			union(e.u, e.v)
			cost += e.weight
		}
	}

	return cost

}
