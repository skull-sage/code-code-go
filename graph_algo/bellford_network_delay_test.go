package graph_algo

import (
	"math"
	"testing"
)

func networkDelayTime(edgeList [][]int, n int, k int) int {
	type Vertex struct {
		id  int
		d   int
		phi *Vertex
	}
	type WEdge struct {
		u *Vertex
		v *Vertex
		w int
	}

	type Graph struct {
		vList    []*Vertex
		edgeList []*WEdge
	}

	init_source := func(g *Graph, s *Vertex) {
		s.d = 0
		s.phi = nil
		for idx := 1; idx < len(g.vList); idx++ {
			v := g.vList[idx]
			if v != s {
				v.d = math.MaxInt
				v.phi = nil
			}

		}
	}

	relax := func(e *WEdge) {
		if e.u.d != math.MaxInt && e.v.d > e.u.d+e.w {
			e.v.d = e.u.d + e.w
			e.v.phi = e.u
		}
	}

	bellman := func(g *Graph, s *Vertex) bool {
		init_source(g, s)

		for idx := 0; idx < n; idx++ {
			for _, edge := range g.edgeList {
				relax(edge)

			}
		}

		for _, edge := range g.edgeList {
			if edge.v.d > edge.u.d+edge.w {
				return false
			}
		}

		return true

	}

	vList := make([]*Vertex, n+1)
	for idx := 1; idx < len(vList); idx++ {
		vList[idx] = &Vertex{
			id: idx,
			d:  math.MaxInt,
		}
	}

	weList := make([]*WEdge, len(edgeList))
	for idx, edge := range edgeList {

		weList[idx] = &WEdge{
			u: vList[edge[0]],
			v: vList[edge[1]],
			w: edge[2],
		}
	}

	g := &Graph{vList: vList, edgeList: weList}
	src := g.vList[k]

	bellman(g, src)

	max := 0
	for idx := 1; idx < len(vList); idx++ {
		v := vList[idx]

		if v != src && v.phi == nil {
			return -1
		} else if v.d > max {
			max = v.d
		}
	}
	return max

}
func TestBFDelay(t *testing.T) {
	times := [][]int{{1, 2, 1}, {2, 1, 3}}
	n := 2
	k := 2
	expected := 3
	actual := networkDelayTime(times, n, k)
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
