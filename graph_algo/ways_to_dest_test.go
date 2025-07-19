package graph_algo

import (
	"container/heap"
	"fmt"
	"math"
	"testing"
)

type qnode struct {
	u     int
	dRank int
}

type PQList []*qnode

func (pq *PQList) Len() int {
	return len(*pq)
}

func (pq *PQList) Swap(i, j int) {
	h := *pq
	h[i], h[j] = h[j], h[i]
}

func (pq *PQList) Less(i, j int) bool {
	h := *pq
	return h[i].dRank > h[j].dRank
}

func (pq *PQList) Push(x any) {
	*pq = append(*pq, x.(*qnode))
}

func (pq *PQList) Pop() any {
	h := *pq
	x := h[len(h)-1]
	*pq = h[0 : len(h)-1]
	return x
}

func countPaths(n int, roads [][]int) int {
	type WEdge struct {
		u int
		v int
		w int
	}

	type AdjList []WEdge
	edgeMap := make(map[int]*AdjList)

	for idx := 0; idx < n; idx++ {
		adjList := new(AdjList)
		edgeMap[idx] = adjList
	}

	for _, edge := range roads {

		u := edge[0]
		v := edge[1]
		w := edge[2]

		uAdj := edgeMap[u]
		vAdj := edgeMap[v]

		*uAdj = append(*uAdj, WEdge{u, v, w})
		*vAdj = append(*vAdj, WEdge{v, u, w})

	}

	dArr := make([]int, n)
	for idx := range dArr {
		dArr[idx] = math.MaxInt
	}
	// dijkstra
	target := n - 1
	pq := new(PQList)
	dArr[target] = 0
	heap.Push(pq, &qnode{target, 0})

	for len(*pq) > 0 {
		uNode := heap.Pop(pq).(*qnode)
		uAdj := edgeMap[uNode.u]

		for _, edge := range *uAdj {

			if dArr[edge.v] > uNode.dRank+edge.w {
				dArr[edge.v] = uNode.dRank + edge.w
				heap.Push(pq, &qnode{edge.v, dArr[edge.v]})
			}
		}
	} // dijkstra ends

	fmt.Println(dArr)
	countPath := make([]int, n)
	_10Pow9 := math.Pow(10, 9)
	_10Pow9Int := int(_10Pow9)

	var dfsCount func(u int) int
	dfsCount = func(u int) int {

		if u == target {
			return 1
		}
		if countPath[u] > 0 {
			return countPath[u]
		}

		adjList := edgeMap[u]
		countU := 0
		for _, edge := range *adjList {
			if dArr[u] == dArr[edge.v]+edge.w {
				countU += dfsCount(edge.v)
				countU = countU % (_10Pow9Int + 7)
			}
		}
		countPath[u] = countU
		return countU
	}

	totalCount := dfsCount(0)
	return totalCount

}

func TestWaySP(t *testing.T) {
	n := 7
	roads := [][]int{
		{0, 6, 7},
		{0, 1, 2},
		{1, 2, 3},
		{1, 3, 3},
		{6, 3, 3},
		{3, 5, 1},
		{6, 5, 1},
		{2, 5, 1},
		{0, 4, 5},
		{4, 6, 2}}

	fmt.Println(countPaths(n, roads))
}
