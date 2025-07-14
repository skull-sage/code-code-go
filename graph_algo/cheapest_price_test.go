package graph_algo

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	type node struct {
		id      int
		dist    int
		adjList []*edge
	}

	type edge struct {
		vNode  *node
		weight int
	}

	nodeMap := make(map[int]*node)
	for idx := 0; idx < n; idx++ {
		nodeMap[idx+1] = &node{id: idx + 1, adjList: make([]*edge, 0)}
	}

	for _, flight := range flights {
		u := nodeMap[flight[0]]
		v := nodeMap[flight[1]]
		u.adjList = append(u.adjList, &edge{vNode: v, weight: flight[2]})
	}

	queue := make([]*node, l, l)

	return 0
}
