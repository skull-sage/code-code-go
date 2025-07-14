package graph_algo

import (
	"fmt"
	"math"
	"testing"
)

type node struct {
	id      int
	adjList []*edge
}

type visit struct {
	id    int
	state int
	dist  int
	step  int
}

type edge struct {
	vNode  *node
	weight int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const (
		NOT_VISITED = iota
		DISCOVERED
		VISITED
	)

	nodeMap := make(map[int]*node)
	visitMap := make(map[int]*visit)
	for idx := 0; idx < n; idx++ {
		nodeMap[idx] = &node{id: idx, adjList: make([]*edge, 0)}
		visitMap[idx] = &visit{id: idx, state: NOT_VISITED, dist: math.MaxInt, step: 0}
	}

	for _, flight := range flights {
		u := nodeMap[flight[0]]
		v := nodeMap[flight[1]]

		u.adjList = append(u.adjList, &edge{vNode: v, weight: flight[2]})
	}

	visitMap[src].dist = 0
	queue := make([]*visit, 0)
	queue = append(queue, visitMap[src])

	for idx := 0; idx < len(queue); idx++ {
		uVisit := queue[idx]
		uVisit.state = DISCOVERED

		for _, edge := range nodeMap[uVisit.id].adjList {
			v := edge.vNode
			vVisit := visitMap[v.id]
			if v.id == dst {
				if vVisit.dist > uVisit.dist+edge.weight {
					vVisit.dist = uVisit.dist + edge.weight
				}
			} else if uVisit.step+1 <= k {

				if vVisit.dist > uVisit.dist+edge.weight {
					vVisit.dist = uVisit.dist + edge.weight
					vVisit.step = uVisit.step + 1
					queue = append(queue, vVisit)
				}
			}
		}
	}

	for _, v := range queue {
		fmt.Printf("(%d %d %d), ", v.id, v.dist, v.step)
	}
	fmt.Println(queue)

	if visitMap[dst].dist != math.MaxInt {
		return visitMap[dst].dist
	} else {
		return -1
	}
}

func TestCheapest(t *testing.T) {

	n := 4
	flights := [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}, {3, 4, 1}}
	src := 0
	dst := 4
	k := 2

	price := findCheapestPrice(n, flights, src, dst, k)
	if price != 7 {
		t.Errorf("expected 7, got %d", price)
	}

	n = 3
	flights = [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src = 0
	dst = 2
	k = 1

	price = findCheapestPrice(n, flights, src, dst, k)
	if price != 200 {
		t.Errorf("expected 200, got %d", price)
	}

	n = 3
	flights = [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src = 0
	dst = 2
	k = 0

	price = findCheapestPrice(n, flights, src, dst, k)
	if price != 500 {
		t.Errorf("expected 500, got %d", price)
	}
}
