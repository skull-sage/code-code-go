package graph

import (
	"fmt"
	"testing"
)

/**
this implementation is still incompelte
need my head around on how to deal with interface in go
*/

const (
	not_discovered = iota
	discovered
	processed
)

type NodeVisit[T comparable] struct {
	phi   T
	dTick int
	fTick int
	state int
}

type dfsVisit[T comparable] struct {
	visitMap  map[T]*NodeVisit[T]
	tickCount int
	tropology []NodeIF[T]
}

func (dfs *dfsVisit[T]) run(u NodeIF[T]) {
	uVisit := dfs.visitMap[u.Id()]
	uVisit.state = discovered
	dfs.tickCount++
	uVisit.dTick = dfs.tickCount

	for _, v := range u.AdjList() {
		vVisit := dfs.visitMap[v.Id()]

		if vVisit.state == not_discovered {
			vVisit.phi = u.Id()
			dfs.run(v)
		}
	}

	dfs.tickCount++
	uVisit.fTick = dfs.tickCount
	uVisit.state = processed
	dfs.tropology = append(dfs.tropology, u)
	fmt.Println("# visit complete: ", u.Id(), "::", uVisit)
	fmt.Println("# ", dfs.visitMap)

}

func NewDFSVisit[T comparable](g graph[T]) dfsVisit[T] {
	visitMap := make(map[T]*NodeVisit[T])
	for id, _ := range g.nodeMap {
		visitMap[id] = &NodeVisit[T]{
			phi:   id,
			state: not_discovered,
		}
	}

	dfs := dfsVisit[T]{visitMap: visitMap, tickCount: 0, tropology: make([]NodeIF[T], 0)}

	for _, u := range g.nodeMap {
		uVisit := visitMap[u.Id()]
		if uVisit.state == not_discovered {
			fmt.Println("# visit start: ", u.Id())
			dfs.run(u)
		}
	}

	return dfs
}

func TestTropology(t *testing.T) {
	g := NewDirectedGraph[string]()
	shirt := &node[string]{id: "shirt"}
	belt := &node[string]{id: "belt"}
	tie := &node[string]{id: "tie"}
	jacket := &node[string]{id: "jacket"}
	socks := &node[string]{id: "socks"}
	undershorts := &node[string]{id: "undershots"}
	pants := &node[string]{id: "pants"}
	shoes := &node[string]{id: "shoes"}
	watch := &node[string]{id: "watch"}

	g.AddNode(shirt)
	g.AddNode(tie)
	g.AddNode(belt)
	g.AddNode(jacket)
	g.AddNode(socks)
	g.AddNode(undershorts)
	g.AddNode(pants)
	g.AddNode(shoes)
	g.AddNode(watch)

	g.AddEdge(shirt, tie)
	g.AddEdge(shirt, belt)
	g.AddEdge(tie, jacket)
	g.AddEdge(belt, jacket)
	g.AddEdge(socks, shoes)
	g.AddEdge(undershorts, shoes)
	g.AddEdge(undershorts, pants)
	g.AddEdge(pants, shoes)
	g.AddEdge(pants, belt)

	dfs := NewDFSVisit(g)
	tropology := dfs.tropology

	fmt.Println("tropology: ")
	for _, node := range tropology {
		fmt.Printf("%s ", node.Id())
		uVisit := dfs.visitMap[node.Id()]
		fmt.Printf("(%d, %d)\n", uVisit.dTick, uVisit.fTick)

	}

}
