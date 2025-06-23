package graph

import (
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
type Node[K comparable] struct {
	Val    any
	AdjSet map[K]bool
}

type Graph[K comparable] struct {
	NodeMap    map[K]*Node[K]
	visitState map[K]int
	extractKey func(any) K
}

func NewGraph[K comparable](extractKey func(any) K) *Graph[K] {
	return &Graph[K]{NodeMap: make(map[K]*Node[K]), visitState: make(map[K]int), extractKey: extractKey}
}

func (self *Graph[K]) String() string {
	str := ""
	for k, v := range self.NodeMap {
		str += fmt.Sprintf("%v: %v\n", k, v)
	}
	return str
}

func (self *Graph[K]) mapNode(val any, k K) *Node[K] {

	node, ok := self.NodeMap[k]
	if !ok {
		node = &Node[K]{Val: val, AdjSet: map[K]bool{}}
		self.NodeMap[k] = node
	}
	return node
}

func (self *Graph[K]) addEdge(u, v any, directed bool) {
	uKey := self.extractKey(u)
	vKey := self.extractKey(v)

	uNode := self.mapNode(u, uKey)
	vNode := self.mapNode(v, vKey)

	uNode.AdjSet[vKey] = true
	if !directed {
		vNode.AdjSet[uKey] = true
	}

}

/*func findMinHeightTrees(n int, edges [][]int) []int {

	if n == 1 {
		return []int{0}
	}

	graph := NewGraph[int](func(a any) int { return a.(int) })

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
}*/

func TestMinHT(t *testing.T) {
	//n := 2
	//edges := [][]int{{0, 1}}

	//roots := findMinHeightTrees(n, edges)
	//fmt.Println("#roots", roots)
	// edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}, {5, 6}}
	// roots := findMinHeightTrees(n, edges)
	// fmt.Println("#roots", roots)
}
