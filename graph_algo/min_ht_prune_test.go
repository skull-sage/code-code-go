package graph_algo

import (
	"container/list"
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
	Key    K
	AdjSet map[K]bool
}

func (self *Node[K]) adjLen() int {
	return len(self.AdjSet)
}

func (self *Node[K]) String() string {
	return fmt.Sprintf("Val: %v, AdjSet: %v", self.Val, self.AdjSet)
}

type Graph[K comparable] struct {
	NodeMap    map[K]*Node[K]
	visitState map[K]int
	extractKey func(any) K
	directed   bool
}

func NewGraph[K comparable](directed bool, extractKey func(any) K) *Graph[K] {
	return &Graph[K]{
		NodeMap:    make(map[K]*Node[K]),
		visitState: make(map[K]int),
		extractKey: extractKey,
		directed:   directed,
	}
}

func (self *Graph[K]) String() string {
	str := ""
	for k, v := range self.NodeMap {
		str += fmt.Sprintf("%v: %v\n", k, v)
	}
	return str
}

func (self *Graph[K]) mapNode(val any, key K) *Node[K] {

	node, ok := self.NodeMap[key]
	if !ok {
		node = &Node[K]{Val: val, AdjSet: map[K]bool{}, Key: key}
		self.NodeMap[key] = node
	}
	return node
}

func (self *Graph[K]) node(k K) *Node[K] {
	node, ok := self.NodeMap[k]
	if !ok {
		panic("node not found")
	}
	return node
}

func (self *Graph[K]) addEdge(u, v any) {
	uKey := self.extractKey(u)
	vKey := self.extractKey(v)
	uNode := self.mapNode(u, uKey)
	vNode := self.mapNode(v, vKey)
	uNode.AdjSet[vKey] = true
	if !self.directed {
		vNode.AdjSet[uKey] = true
	}

}

func (self *Graph[K]) removeEdge(uKey, vKey K) {

	uNode := self.NodeMap[uKey]
	vNode := self.NodeMap[vKey]

	delete(uNode.AdjSet, vKey)
	if !self.directed {
		delete(vNode.AdjSet, uKey)
	}
}

func findMinHeightPrune(n int, edges [][]int) []int {

	if n == 1 {
		return []int{0}
	}

	graph := NewGraph(false, func(a any) int { return a.(int) })

	for idx := range edges {
		edge := edges[idx]
		graph.addEdge(edge[0], edge[1])
	}

	deque := list.New()
	level := 0
	for k, v := range graph.NodeMap {
		if v.adjLen() == 1 {
			deque.PushBack(k)
		}
	}

	remaining := n

	for remaining > 2 {

		qLen := deque.Len()
		remaining = remaining - qLen

		ptr := deque.Front()

		for qLen > 0 {
			uKey := ptr.Value.(int)
			u := graph.node(uKey)

			for vKey := range u.AdjSet {
				adjNode := graph.node(vKey)
				graph.removeEdge(uKey, vKey)
				if adjNode.adjLen() == 1 {
					fmt.Println("Pushing ", vKey)
					deque.PushBack(vKey)
				}
			}

			ptrNext := ptr.Next()
			deque.Remove(ptr)
			qLen--
			ptr = ptrNext

		}

		level++

		// fmt.Printf("#level= %d Deque:", level)
		// for jtr := deque.Front(); jtr != nil; jtr = jtr.Next() {
		// 	fmt.Printf(" %v", jtr.Value.(int))
		// }
		// fmt.Println()

	}

	if remaining == 2 {
		return []int{deque.Front().Value.(int), deque.Back().Value.(int)}
	} else {
		return []int{deque.Front().Value.(int)}
	}

}

func TestMinHT(t *testing.T) {
	//n := 7
	//edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}

	n := 3
	edges := [][]int{{1, 0}, {1, 2}}

	roots := findMinHeightTrees(n, edges)
	fmt.Println("#roots", roots)
}
