package graph_algo

import (
	"fmt"
	"testing"
)

type Node[T comparable] struct {
	Val     T
	AdjList []*Node[T]
}

// NOTE: GO is pass by value-copy, you should always use pointer
func (self *Node[T]) appendAdj(adj *Node[T]) {
	self.AdjList = append(self.AdjList, adj)
}

const (
	not_discovered = iota
	discovered
	visited
)

type Graph[T comparable] struct {
	NodeMap    map[T]*Node[T]
	visitState map[T]int
}

func (self *Graph[T]) addEdge(u, v T, directed bool) {
	uNode, uOK := self.NodeMap[u]
	if !uOK {
		self.NodeMap[u] = &Node[T]{}
	}
	vNode, vOK := self.NodeMap[v]
	if !vOK {
		self.NodeMap[v] = &Node[T]{}
	}
	uNode.appendAdj(vNode)
	if !directed {
		vNode.appendAdj(uNode)
	}

}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{NodeMap: make(map[T]*Node[T]), visitState: make(map[T]int)}
}

func findMinHeightTrees(n int, edges [][]int) []int {
	graph := NewGraph[int]()

	for idx := range edges {
		edge := edges[idx]
		graph.addEdge(edge[0], edge[1], false)
	}

	// inner struct
	type phi struct {
		parent *Node[int]
		length int
	}

	// init phi
	phiArr := make([]phi, n)
	for idx := range phiArr {
		phiArr[idx].parent = nil
		phiArr[idx].length = 0
	}

	// inner function : traverse the input tree
	dfs := func(u *Node[int], phiArr []phi) {
		graph.visitState[u.Val] = discovered

		for _, v := range u.AdjList {
			if graph.visitState[v.Val] == not_discovered {
				phiArr[v.Val].parent = u
				phiArr[v.Val].length = phiArr[u.Val].length + 1
				dfs(v, phiArr)
			}
		}

		graph.visitState[u.Val] = visited

	}

	for _, v := range graph.NodeMap {
		if len(v.AdjList) == 1 {
			dfs(v, phiArr)
		}
	}

	// find the max height
	var maxPhi phi = phiArr[0]
	for idx := 1; idx < len(phiArr); idx++ {
		phi := phiArr[idx]
		if phi.length > maxPhi.length {
			maxPhi = phi
		}
	}

}

func TestMinHT(t *testing.T) {
	n := 6
	edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	roots := findMinHeightTrees(n, edges)
	fmt.Println(roots)
}
