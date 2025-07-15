package graph_algo

import (
	"math"
	"testing"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, kIt int) int {

	cost := make([]int, n)

	for idx := 0; idx < n; idx++ {
		cost[idx] = math.MaxInt
		if idx == src {
			cost[idx] = 0
		}
	}

	temp := make([]int, n)

	for idx := 0; idx < kIt+1; idx++ {

		for idx := 0; idx < n; idx++ {
			temp[idx] = cost[idx]
		}

		for _, e := range flights {
			u := e[0]
			v := e[1]
			uvWeight := e[2]

			if cost[u] != math.MaxInt && temp[v] > cost[u]+uvWeight {
				temp[v] = cost[u] + uvWeight
			}

		}

		for idx := 0; idx < n; idx++ {
			cost[idx] = temp[idx]
		}

	}

	//fmt.Println("src", src, "dst", dst, "cost: ", cost)
	if cost[dst] == math.MaxInt {
		return -1
	}
	return cost[dst]
}

func TestCheapest(t *testing.T) {

	n := 5
	flights := [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}, {3, 4, 1}}
	src := 0
	dst := 4
	kIt := 2

	price := findCheapestPrice(n, flights, src, dst, kIt)
	if price != 7 {
		t.Errorf("expected 7, got %d", price)
	}

	/* n = 3
	flights = [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src = 0
	dst = 2
	kIt = 1

	price = findCheapestPrice(n, flights, src, dst, kIt)
	if price != 200 {
		t.Errorf("expected 200, got %d", price)
	}

	n = 3
	flights = [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src = 0
	dst = 2
	kIt = 0

	price = findCheapestPrice(n, flights, src, dst, kIt)
	if price != 500 {
		t.Errorf("expected 500, got %d", price)
	} */
}
