package graph

type NodeIF[T comparable] interface {
	Id() T
	Weight() int
	AdjList() []NodeIF[T]
	AddAdj(NodeIF[T])
}

type node[T comparable] struct {
	id      T
	weight  int
	adjList []NodeIF[T]
}

func (n *node[T]) Id() T {
	return n.id
}

func (n *node[T]) Weight() int {
	return n.weight
}

func (n *node[T]) AdjList() []NodeIF[T] {
	return n.adjList
}

func (n *node[T]) AddAdj(v NodeIF[T]) {
	n.adjList = append(n.adjList, v)
}

func NewNode[T comparable](id T) NodeIF[T] {
	return &node[T]{
		id: id,
	}
}
