package graph_basics

type Node[T comparable] struct {
	Val    T
	AdjSet map[T]bool
}

// NOTE: GO is pass by value-copy, you should always use pointer
func (self *Node[T]) appendAdj(adj *Node[T]) {
	self.AdjSet[adj.Val] = true
}
