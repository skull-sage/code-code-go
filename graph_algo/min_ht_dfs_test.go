package graph_algo

import (
	"fmt"
	"testing"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := NewGraph[int](false, func(a any) int { return a.(int) })
	for _, edge := range edges {
		graph.addEdge(edge[0], edge[1])
	}

	type phi struct {
		parent     *Node[int]
		length     int
		visitState int
	}

	phiMap := map[int]*phi{}

	var start *Node[int]
	for k, u := range graph.NodeMap {
		phiMap[k] = &phi{
			parent: u,
		}
		if u.adjLen() == 1 {
			start = u
		}
	}

	var dfs func(u *Node[int], length int) int
	dfs = func(u *Node[int], length int) int {
		u.phi.visitState = DISCOVERED

		if u.phi.visitState == NOT_DISCOVERED {
			return u.phi.length
		}
		u.phi.visitState = 1
		u.phi.length = dfs(u.phi.parent) + 1
		return u.phi.length
	}

	return []int{}
}

func TestMhtDFS(t *testing.T) {
	//n := 7
	//edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}

	n := 3
	edges := [][]int{{1, 0}, {1, 2}}

	roots := findMinHeightTrees(n, edges)
	fmt.Println("#roots", roots)
}
