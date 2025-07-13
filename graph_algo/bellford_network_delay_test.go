package graph_algo

import (
	"math"
	"testing"
)

func networkDelayTime(edgeList [][]int, n int, k int) int {
	type Vertex struct {
		d   int
		phi *Vertex
	}
	type Edge struct {
		u *Vertex
		v *Vertex
		w int
	}

	type Graph struct {
		vList    []*Vertex
		edgeList []*Edge
	}

	/* init_source := func(graph *Graph, s *Vertex) {
		s.d = 0
		for _, v := range graph.vList {
			if v == s {
				continue
			}
			v.d = math.MaxInt
			v.phihi = nil
		}
	} */

	relax := func(e *Edge) {
		if e.v.d > e.u.d+e.w {
			e.v.d = e.u.d + e.w
			e.v.phi = e.u
		}
	}

	bellman := func(g *Graph, s *Vertex) bool {
		//init_source(g, s)

		for idx := 0; idx < len(g.vList); idx++ {
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

	g := &Graph{vList: make([]*Vertex, 0), edgeList: make([]*Edge, 0)}

	for _, edge := range edgeList {
		u := &Vertex{
			d: math.MaxInt,
		}
		v := &Vertex{
			d: math.MaxInt,
		}
		e := &Edge{
			u: u,
			v: v,
			w: edge[2],
		}
		g.edgeList = append(g.edgeList, e)
		g.vList = append(g.vList, u)
		g.vList = append(g.vList, v)
	}

	bellman(g, g.vList[k])

	max := 0
	for _, v := range g.vList {
		if v.phi == nil {
			return -1
		} else if v.d > max {
			max = v.d
		}
	}
	return max

}
func TestBFDelay(t *testing.T) {
	times := [][]int{{1, 2, 1}}
	n := 2
	k := 2
	expected := -1
	actual := networkDelayTime(times, n, k)
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
