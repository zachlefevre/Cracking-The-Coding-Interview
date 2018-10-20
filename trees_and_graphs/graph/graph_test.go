package graph

import (
	"fmt"
	"testing"
)

func TestAdd_NewNode_NodeWithNeighbor(t *testing.T) {
	n := GraphNode{value: "root"}
	n.AddNeighbor(GraphNode{value: "firstNeighbor"})
	n.AddNeighbor(GraphNode{value: "secondNeighbor"})
	n.adjacent[0].AddNeighbor(GraphNode{value: "firstNeighborOfFirstNeighbor"})
	fmt.Println(n)
}
