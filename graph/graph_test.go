package graph

import (
	"testing"

	"github.com/google/uuid"
)

func Test_NewNode_NewNeighbor(t *testing.T) {
	idRoot, _ := uuid.NewUUID()
	idN1, _ := uuid.NewUUID()
	r := GraphNode{key: idRoot, value: "root"}
	n1 := GraphNode{key: idN1, value: "neighbor"}
	r.AddNeighbor(n1)

	if r.adjacent[0].key != n1.key {
		t.Fail()
	}
}

func Test_ExistingNode_NewNeighbor(t *testing.T) {
	idRoot, _ := uuid.NewUUID()
	idN1, _ := uuid.NewUUID()
	idN2, _ := uuid.NewUUID()
	r := GraphNode{key: idRoot, value: "root"}
	n1 := GraphNode{key: idN1, value: "neighbor1"}
	n2 := GraphNode{key: idN2, value: "neighbor2"}
	r.AddNeighbor(n1)
	r.AddNeighbor(n2)

	if r.adjacent[0].key != n1.key {
		t.Fail()
	}
	if r.adjacent[1].key != n2.key {
		t.Fail()
	}

}
