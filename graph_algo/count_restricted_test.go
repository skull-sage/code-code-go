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

func countRestrictedPaths(n int, edges [][]int) int {
	type WEdge struct {
		u int
		v int
		w int
	}

	type AdjList []WEdge
	edgeMap := make(map[int]*AdjList)

	for idx := 1; idx <= n; idx++ {
		adjList := new(AdjList)
		edgeMap[idx] = adjList
	}

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]
		w := edge[2]

		uAdj := edgeMap[u]
		vAdj := edgeMap[v]

		*uAdj = append(*uAdj, WEdge{u, v, w})
		*vAdj = append(*vAdj, WEdge{v, u, w})

	}
	type Dist struct {
		d   int
		phi int
	}
	dArr := make([]*Dist, n+1)
	for idx := range dArr {
		dArr[idx] = &Dist{d: math.MaxInt, phi: -1}
	}
	// dijkstra
	start := n
	pq := new(PQList)
	dArr[start] = &Dist{d: 0, phi: start}
	heap.Push(pq, &qnode{start, 0})

	for len(*pq) > 0 {
		uNode := heap.Pop(pq).(*qnode)
		uAdj := edgeMap[uNode.u]

		for _, edge := range *uAdj {
			distV := dArr[edge.v] // we will use ptr ref distV

			if distV.d > uNode.rank+edge.w {
				distV.d = uNode.rank + edge.w
				distV.phi = uNode.u
				heap.Push(pq, &qnode{edge.v, distV.d})
			}
		}
	}

	edgeMap = make(map[int]*AdjList)
	for _, e := range edges {
		u := e[0]
		v := e[1]
		if u < v && dArr[u].d > dArr[v].d {
			adjU := edgeMap[u]
			*adjU = append(*adjU, WEdge{u, v, 0}) // we con't care about w
		} else if dArr[v].d > dArr[u].d {
			adjV := edgeMap[v]
			*adjV = append(*adjV, WEdge{v, u, 0})
		}
	}

	var dfsCount func(u int) int
	dfsCount = func(u int) int {
		// TO BE Implemented
		return 0
	}

	total := dfsCount(1)
	_10Pow9 := math.Pow(10, 9)
	_10Pow9Int := int(_10Pow9)
	return total % (_10Pow9Int + 7)

	//fmt.Println("# u", uNode)
	//fmt.Println("#heap", *pq)
	//fmt.Println("#dArr", dArr)

}
