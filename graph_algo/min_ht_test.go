package graph_algo

import (
	graph "algo/main/graph"
	"fmt"
	"testing"
)

const (
	NOT_DISCOVERED = iota
	DISCOVERED
	VISITED
)

// T is the generic type of id::identity for graph nodes

// NOTE: GO is pass by value-copy, you should always use pointer
type Node[T comparable] struct {
	Val    T
	AdjSet map[T]bool
}

type Graph[T comparable] struct {
	NodeMap    map[T]*graph.Node[T]
	visitState map[T]int
}

func (self *Graph[T]) mapNode(val T) *graph.Node[T] {
	node, ok := self.NodeMap[val]
	if !ok {
		node = &Node[T]{Val: val, AdjSet: map[T]bool{}}
		self.NodeMap[val] = node
	}
	return node
}

func (self *Graph[T]) String() string {
	str := ""
	for k, v := range self.NodeMap {
		str += fmt.Sprintf("%v: %v\n", k, v)
	}
	return str
}

func (self *Graph[T]) addEdge(u, v T, directed bool) {
	uNode := self.mapNode(u)
	vNode := self.mapNode(v)

	uNode.appendAdj(vNode)
	if !directed {
		vNode.appendAdj(uNode)
	}

}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{NodeMap: make(map[T]*Node[T]), visitState: make(map[T]int)}
}

func findMinHeightTrees(n int, edges [][]int) []int {

	if n == 1 {
		return []int{0}
	}

	graph := NewGraph[int]()

	for idx := range edges {
		edge := edges[idx]
		graph.addEdge(edge[0], edge[1], false)
	}

	//fmt.Println("# Created Graph: ", graph)

	// inner struct
	type phi struct {
		node   *Node[int]
		parent *Node[int]
		length int
	}

	// init phi
	phiArr := make([]phi, n)
	for idx := range phiArr {
		phiArr[idx].node = graph.NodeMap[idx]
		phiArr[idx].parent = nil
		phiArr[idx].length = 0
	}

	for k := range graph.NodeMap {
		graph.visitState[k] = NOT_DISCOVERED
	}

	var dfs func(u *Node[int], phiArr []phi)
	// inner function : traverse the input tree
	dfs = func(u *Node[int], phiArr []phi) {
		graph.visitState[u.Val] = DISCOVERED

		for k, _ := range u.AdjSet {
			v := graph.NodeMap[k]
			if graph.visitState[v.Val] == NOT_DISCOVERED {
				phiArr[v.Val].parent = u
				phiArr[v.Val].length = phiArr[k].length + 1
				dfs(v, phiArr)
			}
		}

		graph.visitState[u.Val] = VISITED

	}

	for _, v := range graph.NodeMap {
		if len(v.AdjSet) == 1 {

			//fmt.Println("# Start DFS with v: ", v.Val)
			dfs(v, phiArr)
			break
		}
	}

	// fmt.Println("# Calculated phi: ")
	// for idx := range phiArr {
	// 	fmt.Printf("%d->len:%d ", phiArr[idx].node.Val, phiArr[idx].length)
	// }

	// find the max height
	var maxPhi phi = phiArr[0]
	for idx := 1; idx < len(phiArr); idx++ {
		phi := phiArr[idx]
		if phi.length > maxPhi.length {
			maxPhi = phi
		}
	}

	//fmt.Printf("\n# Max phi: %d->len:%d\n", maxPhi.node.Val, maxPhi.length)

	traceBack := maxPhi.length / 2
	aPhi := maxPhi
	for idx := traceBack; idx >= 1; idx-- {
		aPhi = phiArr[aPhi.parent.Val]
		//fmt.Printf("%d->len:%d ", aPhi.node.Val, aPhi.length)
	}

	//fmt.Println()
	// total number of nodes = maxPhi.length + 1
	if (maxPhi.length+1)%2 == 0 {
		return []int{aPhi.node.Val, aPhi.parent.Val}
	} else {
		return []int{aPhi.node.Val}
	}
}

func TestMinHT(t *testing.T) {
	n := 2
	edges := [][]int{{0, 1}}
	roots := findMinHeightTrees(n, edges)
	fmt.Println("#roots", roots)
	// edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}, {5, 6}}
	// roots := findMinHeightTrees(n, edges)
	// fmt.Println("#roots", roots)
}
