package graph_algo

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

func findMhTByPruning(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := NewGraph[int](false, func(a any) int { return a.(int) })
	for _, edge := range edges {
		graph.addEdge(edge[0], edge[1])
	}

	levelMap := make(map[int]int)

	queue := make([]*Node[int], 0)
	for _, u := range graph.NodeMap {
		if u.adjLen() == 1 {
			queue = append(queue, u)
			levelMap[u.Key] = 1
		}
	}

	height := 1
	for idx := 0; idx < len(queue); idx++ {
		u := queue[idx] // u is a leaf node
		uLevel := levelMap[u.Key]

		if height < uLevel {
			height = uLevel
		}

		for vKey := range u.AdjSet { // this loop only run only once as u is lead node means u.adjLen() == 1
			v := graph.node(vKey)
			vLevel, ok := levelMap[vKey]

			if ok && vLevel == uLevel && v.adjLen() == 1 {
				// as v is now a leaf node with same level to u, (u, v) is the top
				// hence our level traversal ends
				break
			} else {

				if !ok || vLevel < uLevel+1 {
					levelMap[vKey] = uLevel + 1
				}
				graph.removeEdge(vKey, u.Key)
				if v.adjLen() == 1 {
					queue = append(queue, v)
				}
			}

		}

		//fmt.Println("#levelMap", levelMap)
	}

	result := make([]int, 0, 0)
	for key, level := range levelMap {
		if level == height {
			result = append(result, key)
		}
	}

	return result
}

func TestMinHT(t *testing.T) {
	//n := 7
	//edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}

	n := 3
	edges := [][]int{{1, 0}, {1, 2}}

	roots := findMhTByPruning(n, edges)
	fmt.Println("#roots", roots)
}
