package graph_algo

import (
	"container/heap"
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

func countRestrictedPaths(n int, edges [][]int) int {
	type WEdge struct {
		u int
		v int
		w int
	}

	type AdjList []WEdge
	nodeEdgeList := make([]*AdjList, n+1, n+1)

	for idx := 1; idx <= n; idx++ {
		adjList := new(AdjList)
		nodeEdgeList[idx] = adjList
	}

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]
		w := edge[2]

		uAdj := nodeEdgeList[u]
		vAdj := nodeEdgeList[v]

		*uAdj = append(*uAdj, WEdge{u, v, w})
		*vAdj = append(*vAdj, WEdge{v, u, w})

	}

	dArr := make([]int, n+1)
	for idx := range dArr {
		dArr[idx] = math.MaxInt
	}
	// dijkstra
	target := n
	pq := new(PQList)
	dArr[target] = 0
	heap.Push(pq, &qnode{target, 0})

	for len(*pq) > 0 {
		uNode := heap.Pop(pq).(*qnode)
		uAdj := nodeEdgeList[uNode.u]

		for _, edge := range *uAdj {

			if dArr[edge.v] > uNode.dRank+edge.w {
				dArr[edge.v] = uNode.dRank + edge.w
				heap.Push(pq, &qnode{edge.v, dArr[edge.v]})
			}
		}
	}

	countPath := make([]int, n+1)
	_10Pow9 := math.Pow(10, 9)
	_10Pow9Int := int(_10Pow9)

	var dfsCount func(u int) int
	dfsCount = func(u int) int {

		/* if u == target {
			return 1
		}
		if countPath[u] > 0 {
			return countPath[u]
		} */

		adjList := nodeEdgeList[u]
		countU := 0
		for _, edge := range *adjList {
			if edge.v == target {
				countU += 1
			} else if dArr[u] > dArr[edge.v] {
				if countPath[edge.v] > 0 {
					countU += countPath[edge.v]
				} else {
					countU += dfsCount(edge.v)
				}
				countU = countU % (_10Pow9Int + 7)
			}
		}
		countPath[u] = countU
		return countU
	}

	totalCount := dfsCount(1)
	return totalCount

	//fmt.Println("# u", uNode)
	//fmt.Println("#heap", *pq)
	//fmt.Println("#dArr", dArr)

}

func TestRestricted(t *testing.T) {

	// restricted cases are removed for future ref
	// the solution face constant TLE on Leetcode for no apperant reason
}
