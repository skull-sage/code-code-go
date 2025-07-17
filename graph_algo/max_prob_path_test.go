package graph_algo

import (
	"fmt"
	"testing"
)

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	type WEdge struct {
		u int
		v int
		w float64
	}

	wList := make([][]WEdge, len(edges))

	for idx, edge := range edges {
		adjMap[edge[0]] = &WEdge{
			u: edge[0],
			v: edge[1],
			w: succProb[idx],
		}
		adjMap[edge[1]] = &WEdge{
			u: edge[1],
			v: edge[0],
			w: succProb[idx],
		}
	}

	dArr := make([]float64, n)
	//bellman-ford init
	dArr[start] = 1.0
	for idx := 1; idx < n; idx++ {
		dArr[idx] = 0 //math.MaxFloat64
	}
	// init ends

	// bellman-relax
	for idx := 0; idx < n; idx++ {
		for _, wedge := range adjMap {
			if dArr[wedge.v] < dArr[wedge.u]*wedge.w {
				dArr[wedge.v] = dArr[wedge.u] * wedge.w
			}
			fmt.Println("#", idx, wedge, "=>", dArr)
		}
		fmt.Println(dArr)
	} // relax ends

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
