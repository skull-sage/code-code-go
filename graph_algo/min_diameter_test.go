package graph_algo

import (
	"fmt"
	"math"
	"testing"
)

func findDiameter(edges [][]int) (int, int) {
	n := len(edges)
	if n == 0 {
		return 0, 0
	} else if n == 1 {
		return 1, 1
	} else if n == 2 {
		return 1, 2
	}

	graph := NewGraph[int](false, func(a any) int { return a.(int) })
	for _, edge := range edges {
		graph.addEdge(edge[0], edge[1])
	}

	levelMap := make(map[int]int)

	queue := make([]*Node[int], 0)
	for _, u := range graph.NodeMap {
		if u.adjLen() == 1 {
			queue = append(queue, u)
			levelMap[u.Key] = 0
		}
	}

	height := 0
	for idx := 0; idx < len(queue); idx++ {
		u := queue[idx] // u is a leaf node
		uLevel := levelMap[u.Key]

		if height < uLevel {
			height = uLevel
		}

		for vKey := range u.AdjSet { // this loop only run only once as u is lead node means u.adjLen() == 1
			v := graph.node(vKey)
			vLevel, ok := levelMap[vKey]

			if ok && vLevel == uLevel && v.adjLen() == 1 {
				// as v is now a leaf node with same level to u, (u, v) is the top
				// hence our level traversal ends
				break
			} else {

				if !ok || vLevel < uLevel+1 {
					levelMap[vKey] = uLevel + 1
				}
				graph.removeEdge(vKey, u.Key)
				if v.adjLen() == 1 {
					queue = append(queue, v)
				}
			}

		}

		//fmt.Println("#levelMap", levelMap)
	}

	result := make([]int, 0, 0)
	for key, level := range levelMap {
		if level == height {
			result = append(result, key)
		}
	}

	if len(result) == 2 {
		return height, 2*height + 1
	} else {
		return height, 2 * height
	}
}

func findMergedDiameter(edges1 [][]int, edges2 [][]int) int {
	_, d1 := findDiameter(edges1)
	_, d2 := findDiameter(edges2)

	//fmt.Println("h1", h1, "h2", h2)
	//fmt.Println("d1", d1, "d2", d2)
	combined := int(math.Ceil(float64(d1)/2.0)) + int(math.Ceil(float64(d2)/2.0)) + 1
	//fmt.Println("combined", combined)
	result := d1
	if result < d2 {
		result = d2
	}

	if result < combined {
		result = combined
	}

	return result
}

func TestMinDiameter(t *testing.T) {
	//n := 7
	//edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}

	//edges1 := [][]int{{1, 0}, {0, 2}, {0, 3}}
	//edges2 := [][]int{{0, 1}}

	edges1 := [][]int{{1, 0}}
	edges2 := [][]int{{0, 1}}

	result := findMergedDiameter(edges1, edges2)
	fmt.Println("#result", result)
}
