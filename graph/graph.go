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

func TestMinHT(t *testing.T) {
	//n := 2
	//edges := [][]int{{0, 1}}

	//roots := findMinHeightTrees(n, edges)
	//fmt.Println("#roots", roots)
	// edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}, {5, 6}}
	// roots := findMinHeightTrees(n, edges)
	// fmt.Println("#roots", roots)
}
