package graph

type GraphIF[T comparable] interface {
	AdjList(nodeId T) []NodeIF[T]
	AddEdge(u, v NodeIF[T])
	AddNode(u NodeIF[T])
	IsDirected() bool
}

type graph[T comparable] struct {
	nodeMap  map[T]NodeIF[T]
	directed bool
}

func (g *graph[T]) AdjList(nodeId T) []NodeIF[T] {
	var n NodeIF[T]
	n = g.nodeMap[nodeId]
	return n.AdjList()
}

func (g *graph[T]) IsDirected() bool {
	return g.directed
}

func (g *graph[T]) AddNode(u NodeIF[T]) {
	g.nodeMap[u.Id()] = u
}

func (g *graph[T]) AddEdge(u, v NodeIF[T]) {
	uNode := g.nodeMap[u.Id()]
	vNode := g.nodeMap[v.Id()]
	if g.IsDirected() {
		uNode.AddAdj(vNode)
	} else {
		uNode.AddAdj(vNode)
		vNode.AddAdj(uNode)
	}
}

func NewGraph[T comparable]() graph[T] {
	return graph[T]{
		nodeMap:  make(map[T]NodeIF[T]),
		directed: false,
	}
}

func NewDirectedGraph[T comparable]() graph[T] {
	return graph[T]{
		nodeMap:  make(map[T]NodeIF[T]),
		directed: true,
	}
}
