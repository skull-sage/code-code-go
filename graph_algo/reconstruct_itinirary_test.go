package graph_algo

import (
	"fmt"
	"sort"
	"testing"
)

func findEdgePath(edgeMap map[string][]string, start string, path []string) []string {

	fmt.Printf("start: %s, path: %v\n", start, path)

	for len(edgeMap[start]) > 0 {
		next := edgeMap[start][0]           // we pick an edge
		edgeMap[start] = edgeMap[start][1:] // we delete the picked edge
		path = findEdgePath(edgeMap, next, path)
	}

	return append(path, start)

}

func findItinerary(tickets [][]string) []string {

	edgeMap := make(map[string][]string)

	for _, ticket := range tickets {
		edgeMap[ticket[0]] = append(edgeMap[ticket[0]], ticket[1])
	}

	for _, edgeList := range edgeMap {
		sort.Strings(edgeList)
	}

	path := findEdgePath(edgeMap, "JFK", make([]string, 0))

	revPath := make([]string, 0, len(path))

	// reverse the path
	for i := len(path) - 1; i >= 0; i-- {
		revPath = append(revPath, path[i])
	}

	fmt.Println("path:", path)
	fmt.Println("reverse:", revPath)

	return revPath

}

func TestFindItinerary(t *testing.T) {
	tickets := [][]string{
		{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "JFK"}, {"ATL", "AAA"}, {"AAA", "ATL"},
		//,
	}

	findItinerary(tickets)

	tickets = [][]string{
		{"JFK", "SFO"}, {"JFK", "ATL"}, {"ATL", "JFK"}, {"SFO", "AAA"}, {"AAA", "SFO"}, {"SFO", "JFK"},
	}

	findItinerary(tickets)

}
