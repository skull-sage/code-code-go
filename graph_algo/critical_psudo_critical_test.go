package graph_algo

type vertex struct {
	 
	rank int
	p    *vertex
}

func link(x, y *vertex) {
	if x.rank > y.rank {
		y.p = x
	} else {
		x.p = y
		if x.rank == y.rank {
			y.rank = y.rank + 1
		}
	}
}

func findSet(v *vertex) *vertex {
	if v != v.p {
		v.p = findSet(v.p)
	}
	return v.p
}

func isDisjoint(u, v *vertex) bool {
	return findSet(u) != findSet(v)
}

func union(x, y *vertex) {
	px := findSet(x)
	py := findSet(y)
	link(px, py)
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {

}
