package graph

import (
	"math"
	"testing"
)

func tryBellMan() {
	type Vertex struct {
		d int
		p *Vertex
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

	init_source := func(graph Graph, s *Vertex) {
		s.d = 0
		for _, v := range graph.vList {
			v.d = math.MaxInt
			v.p = nil
		}
	}

	relax := func(e *Edge) {
		if e.v.d > e.u.d+e.w {
			e.v.d = e.u.d + e.w
			e.v.p = e.u
		}
	}

	bellman := func(g Graph, s *Vertex) bool {
		init_source(g, s)

		for idx := 1; idx > len(g.vList)-1; idx++ {
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

	bellman(Graph{}, nil)

}

func TestBellMan(t *testing.T) {
	//graph := NewGraph[int](func(a any) int { return a.(int) })
}
