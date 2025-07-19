package graph_algo

import (
	"container/heap"
	"math"
)

type qnode struct {
	u    int
	rank int
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
	return h[i].rank > h[j].rank
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

	dArr := make([]int, n+1)
	for idx := range dArr {
		dArr[idx] = math.MaxInt
	}
	// dijkstra
	start := n
	pq := new(PQList)

	heap.Push(pq, &qnode{start, 0})

	for len(*pq) > 0 {
		uNode := heap.Pop(pq).(*qnode)
		uAdj := edgeMap[uNode.u]

		for _, edge := range *uAdj {

			if dArr[edge.v] > uNode.dRank+edge.w {
				distV.d = uNode.rank + edge.w
				distV.phi = uNode.u
				heap.Push(pq, &qnode{edge.v, distV.d})
			}
		}
	}

	dest := n - 1

}
