package graph_algo

/*
const (
	not_discovered = iota
	discovered
	processed
)

type WEdge struct {
	weight int
	vId    int
}

func (w WEdge) String() string {
	return fmt.Sprintf("weight: %d, vId: %d", w.weight, w.vId)
}

type node struct {
	id      int
	d       int
	phi     *node
	adjList []*WEdge
}

func (n node) String() string {
	return fmt.Sprintf("id: %d, d: %d, adj: %v", n.id, n.d, n.adjList)
}

func (n *node) AddAdj(e *WEdge) {
	n.adjList = append(n.adjList, e)
}

type graph struct {
	nodeMap  map[int]*node
	directed bool
}

func (g *graph) AdjList(nodeId int) []*WEdge {
	n := g.nodeMap[nodeId]
	return n.adjList
}

func (g *graph) IsDirected() bool {
	return g.directed
}

func (g *graph) AddNode(u *node) {
	g.nodeMap[u.id] = u
}

func (g *graph) AddEdge(uId, vId int, weight int) {
	uNode := g.nodeMap[uId]
	uNode.AddAdj(&WEdge{weight: weight, vId: vId})
}

func NewNetworkGraph() graph {
	return graph{
		nodeMap:  make(map[int]*node),
		directed: false,
	}
}

type NodeVisit struct {
	phi   int
	dTick int
	fTick int
	state int
}

type dfsVisit struct {
	g         *graph
	visitMap  map[int]*NodeVisit
	tickCount int
	tropology []*node
}

func (dfs *dfsVisit) run(uId int) {
	u := dfs.g.nodeMap[uId]
	uVisit := dfs.visitMap[uId]
	uVisit.state = discovered
	dfs.tickCount++
	uVisit.dTick = dfs.tickCount

	for _, edge := range u.adjList {
		vVisit := dfs.visitMap[edge.vId]

		if vVisit.state == not_discovered {
			vVisit.phi = u.id
			dfs.run(edge.vId)
		}
	}

	dfs.tickCount++
	uVisit.fTick = dfs.tickCount
	uVisit.state = processed
	dfs.tropology = append(dfs.tropology, u)

}

func NewDFSVisit(g *graph) *dfsVisit {

	visitMap := make(map[int]*NodeVisit)
	for id, _ := range g.nodeMap {
		visitMap[id] = &NodeVisit{
			phi:   id,
			state: not_discovered,
		}
	}

	dfs := &dfsVisit{g: g, visitMap: visitMap, tickCount: 0, tropology: make([]*node, 0)}

	return dfs
}

func buildGraph(times [][]int, n int) *graph {
	g := NewNetworkGraph()
	for idx := 0; idx < n; idx++ {
		g.AddNode(&node{id: idx + 1, adjList: make([]*WEdge, 0)})
	}
	for _, time := range times {
		g.AddEdge(time[0], time[1], time[2])
	}
	return &g
}

func networkDelayTime_dag(times [][]int, n int, k int) int {
	g := buildGraph(times, n)

	dfs := NewDFSVisit(g)
	dfs.run(k)

	if len(dfs.tropology) < len(g.nodeMap) {
		return -1
	}

	// initialize
	sNode := g.nodeMap[k]
	sNode.d = 0
	sNode.phi = sNode
	for _, uNode := range g.nodeMap {
		if uNode != sNode {
			uNode.d = math.MaxInt
			uNode.phi = nil
		}

	}

	for idx := 0; idx < len(dfs.tropology); idx++ {
		u := dfs.tropology[idx]
		//fmt.Println("# u:", u)

		//fmt.Println("# ", u)
		for _, uvEdge := range u.adjList {
			v := g.nodeMap[uvEdge.vId]
			weight := uvEdge.weight

			if v.d > u.d+weight {
				v.d = u.d + weight
				v.phi = u
			}

		}

	}

	max := 0
	for _, u := range dfs.tropology {
		//fmt.Println("#u:", u, u.phi)

		if u.phi == nil {
			return -1
		} else if u.d > max {
			max = u.d
		}
	}
	return max

}

func TestNetworkDelayTime(t *testing.T) {
	times := [][]int{{1, 2, 1}}
	n := 2
	k := 2
	expected := -1
	actual := networkDelayTime_dag(times, n, k)
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
*/
