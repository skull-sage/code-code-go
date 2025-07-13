package graph_algo

import (
	"testing"
)

const (
	not_discovered = iota
	discovered
	processed
)

type WEdge struct {
	weight int
	vId    int
}

type node struct {
	id      int
	adjList []*WEdge
}

func (n *node) AddAdj(e *WEdge) {
	n.adjList = append(n.adjList, e)
}

type graph struct {
	nodeMap  map[int]*node
	directed bool
}

func (g *graph) AdjList(nodeId int) []*WEdge {
	n := g.nodeMap[nodeId]
	return n.adjList
}

func (g *graph) IsDirected() bool {
	return g.directed
}

func (g *graph) AddNode(u *node) {
	g.nodeMap[u.id] = u
}

func (g *graph) AddEdge(uId, vId int, weight int) {
	uNode := g.nodeMap[uId]
	uNode.AddAdj(&WEdge{weight: weight, vId: vId})
}

func NewNetworkGraph() graph {
	return graph{
		nodeMap:  make(map[int]*node),
		directed: false,
	}
}

type NodeVisit struct {
	phi   int
	dTick int
	fTick int
	state int
}

type dfsVisit struct {
	g         graph
	visitMap  map[int]*NodeVisit
	tickCount int
	tropology []*node
}

func (dfs *dfsVisit) run(uId int) {
	u := dfs.g.nodeMap[uId]
	uVisit := dfs.visitMap[uId]
	uVisit.state = discovered
	dfs.tickCount++
	uVisit.dTick = dfs.tickCount

	for _, edge := range u.adjList {
		vVisit := dfs.visitMap[edge.vId]

		if vVisit.state == not_discovered {
			vVisit.phi = u.id
			dfs.run(edge.vId)
		}
	}

	dfs.tickCount++
	uVisit.fTick = dfs.tickCount
	uVisit.state = processed
	dfs.tropology = append(dfs.tropology, u)

}

func NewDFSVisit(g graph) dfsVisit {
	visitMap := make(map[int]*NodeVisit)
	for id, _ := range g.nodeMap {
		visitMap[id] = &NodeVisit{
			phi:   id,
			state: not_discovered,
		}
	}

	dfs := dfsVisit{visitMap: visitMap, tickCount: 0, tropology: make([]*node, 0)}

	for uId, _ := range g.nodeMap {
		uVisit := visitMap[uId]
		if uVisit.state == not_discovered {
			dfs.run(uId)
		}
	}

	return dfs
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := NewNetworkGraph()
	for idx := 0; idx < n; idx++ {
		graph.AddNode(&node{id: idx + 1, adjList: make([]*WEdge, 0)})
	}
	for _, time := range times {
		graph.AddEdge(time[0], time[1], time[2])
	}
}

func TestNetworkDelayTime(t *testing.T) {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 2
	expected := 2
	actual := networkDelayTime(times, n, k)
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
