package graph_algo

import (
	"testing"
)

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	type Edge struct {
		v       int
		wCount int
		adjList []*Edge
	}

	type Visit struct {
		u     int
		count int
	}

	edgeMap := make([]*Edge, n)
	for idx := 0; idx < n; idx++ {
		edgeMap[idx] = &Edge{v: idx, adjList: make([]*Edge, 0)}
	}

	for _, e := range edges {
			u := edgeMap[e[0]]
			u.adjList = append(u.adjList, &Edge{v: e[1], wCount: e[2]+1})
	}

	queue := make([]*Visit, 0)
	queue = append(queue, &Visit{u: 0, count: 1})

	totalCount :=0 
 	for idx:=0; idx< n; idx++ {
		head := queue[idx]
		 
		adjList := edgeMap[head.u].adjList
		for _, edge := range adjList {
			 if head.count + edge.wCount <= maxMoves {
				totalCount += edge.wCount
				queue = append(queue, &Visit{u: edge.v, count: head.count + edge.wCount})
			 }
		}
}

func TestSubdividedReachable(t *testing.T) {

}
