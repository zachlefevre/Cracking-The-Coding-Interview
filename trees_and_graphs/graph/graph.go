package graph

type GraphNode struct {
	value    interface{}
	adjacent []GraphNode
}

func (n *GraphNode) AddNeighbor(a GraphNode) {
	n.adjacent = append(n.adjacent, a)
}

func (n GraphNode) FindPathDFS(d GraphNode) []GraphNode {
	visited := make(map[*GraphNode]bool)
	var s Stack
	s = s.push(n)
	for len(s) != 0 {

	}
	return []GraphNode{}
}
