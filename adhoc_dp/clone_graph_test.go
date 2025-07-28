package adhoc_dp

import (
	"testing"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

var visitMap map[int]*Node

func createClone(node *Node) *Node {
	return &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
}

func cloneNode(node *Node) *Node {

	newNode := createClone(node)
	visitMap[node.Val] = newNode

	for idx, u := range node.Neighbors {
		cloneU, ok := visitMap[u.Val]
		if !ok {
			cloneU = cloneNode(u)
		}
		newNode.Neighbors[idx] = cloneU
	}

	return newNode
}

func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}
	visitMap = make(map[int]*Node)
	newClone := cloneNode(node)

	return newClone

}

func TestCloneGraph(t *testing.T) {

}
