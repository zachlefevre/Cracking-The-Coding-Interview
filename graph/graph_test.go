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

func TestDFS_ExistingGraph_Path(t *testing.T) {
	n := GraphNode{key: 0, value: "a"}
	n.AddNeighbor(GraphNode{key: 01, value: "b"})
	n.AddNeighbor(GraphNode{key: 02, value: "c"})
	n.adjacent[0].AddNeighbor(GraphNode{key: 03, value: "d"})
	n.adjacent[1].AddNeighbor(GraphNode{key: 04, value: "e"})
	n.adjacent[1].adjacent[0].AddNeighbor(GraphNode{key: 05, value: "f"})
	path, _ := n.FindPathDFS(03)
	for _, p := range path {
		fmt.Println(p.key)
	}
	fmt.Println(n)
}

func TestBFS_ExistingGraph_Path(t *testing.T) {
	n := GraphNode{key: 0, value: "a"}
	n.AddNeighbor(GraphNode{key: 01, value: "b"})
	n.AddNeighbor(GraphNode{key: 02, value: "c"})
	n.AddNeighbor(GraphNode{key: 06, value: "g"})
	n.adjacent[0].AddNeighbor(GraphNode{key: 03, value: "d"})
	n.adjacent[1].AddNeighbor(GraphNode{key: 04, value: "e"})
	n.adjacent[1].adjacent[0].AddNeighbor(GraphNode{key: 05, value: "f"})
	path, _ := n.FindPathBFS(03)
	for _, p := range path {
		fmt.Println(p.key)
	}
	fmt.Println(n)
}
