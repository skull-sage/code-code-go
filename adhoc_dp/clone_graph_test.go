package adhoc_dp

type Node struct {
	Val       int
	Neighbors []*Node
}

var visitMap map[int]*Node

func createClone(node *Node) *Node {
	return &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
}

func cloneNode(node *Node) *Node {

	newNode := createClone(node)
	visitMap[node.Val] = newNode

	for _, u := range node.Neighbors {
		_, ok := visitMap[u.Val]
		if !ok {
			cloneU := cloneNode(u)
			newNode.Neighbors = append(newNode.Neighbors, cloneU)
			cloneU.Neighbors = append(cloneU.Neighbors, newNode)
		}

	}

	return newNode
}

func cloneGraph(node *Node) *Node {

	visitMap = make(map[int]*Node)
	return cloneNode(node)

}
