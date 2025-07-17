package graph_algo

import (
	"fmt"
	"math"
	"testing"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// to be implemented

	dist := make([][]int, n, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			dist[i][j] = math.MaxInt
		}
	}
	for _, edge := range edges {
		dist[edge[0]][edge[1]] = edge[2]
		dist[edge[1]][edge[0]] = edge[2]
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] == math.MaxInt || dist[k][j] == math.MaxInt {
					continue
				}
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	rCount := make([]int, n, n)
	for i := 0; i < n; i++ {
		//fmt.Printf("for %d => ", i)
		for j := 0; j < n; j++ {
			//fmt.Printf("%d ", dist[i][j])
			if dist[i][j] <= distanceThreshold && i != j {
				rCount[i]++
			}
		}
		//fmt.Println()
	}

	gtMinCity := 0

	for i := 0; i < n; i++ {
		//fmt.Println("# ", i, "=>", rCount[i])
		if rCount[gtMinCity] > rCount[i] {
			gtMinCity = i
		} else if rCount[gtMinCity] == rCount[i] {
			gtMinCity = max(gtMinCity, i)
		}

	}
	return gtMinCity
}

func TestFindTheCity(t *testing.T) {
	// to be implemented
	n := 4
	edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
	distanceThreshold := 4
	gtMinCity := findTheCity(n, edges, distanceThreshold)
	fmt.Println(gtMinCity)
}
