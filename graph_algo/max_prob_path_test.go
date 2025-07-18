package graph_algo

import (
	"container/heap"
	"fmt"
	"testing"
)

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {

	type WEdge struct {
		u int
		v int
		w float64
	}

	type AdjList []WEdge
	edgeMap := make(map[int]*AdjList)

	for idx := 0; idx < n; idx++ {
		adjList := new(AdjList)
		edgeMap[idx] = adjList
	}

	for idx, edge := range edges {

		u := edge[0]
		v := edge[1]
		w := succProb[idx]

		uAdj := edgeMap[u]
		vAdj := edgeMap[v]

		*uAdj = append(*uAdj, WEdge{u, v, w})
		*vAdj = append(*vAdj, WEdge{v, u, w})

	}

	dArr := make([]float64, n)
	// dijkstra

	pq := new(PQList)
	dArr[start] = 1
	heap.Push(pq, &qnode{start, 1})
	for len(*pq) > 0 {
		uNode := heap.Pop(pq).(*qnode)
		uAdj := edgeMap[uNode.u]

		for _, edge := range *uAdj {
			fmt.Println(edge)
			/* if dArr[edge.v] < uNode.rank*edge.w {
				dArr[edge.v] = dArr[edge.u] * edge.w
				heap.Push(pq, &qnode{edge.v, dArr[edge.v]})
			} */
		}

		//fmt.Println("# u", uNode)
		//fmt.Println("#heap", *pq)
		//fmt.Println("#dArr", dArr)
	}

	return dArr[end] // relax ends

}

func TestMaxProb(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
	succProb := []float64{0.5, 0.5, 0.2}
	start := 0
	end := 2
	t.Log(maxProbability(n, edges, succProb, start, end))
}
