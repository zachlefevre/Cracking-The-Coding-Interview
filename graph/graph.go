package graph

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/zachlefevre/Cracking-The-Coding-Interview/queue"
	"github.com/zachlefevre/Cracking-The-Coding-Interview/stack"
)

type GraphNode struct {
	key      uuid.UUID
	value    interface{}
	adjacent []GraphNode
}

func (n *GraphNode) AddNeighbor(a GraphNode) {
	n.adjacent = append(n.adjacent, a)
}

func (n GraphNode) FindPathDFS(goal uuid.UUID) (path []GraphNode, e error) {
	visited := make(map[uuid.UUID]bool)
	var s stack.Stack
	s = s.Push(n)
	for {
		if s.IsEmpty() {
			break
		}
		var err error
		var val interface{}
		s, val, err = s.Pop()
		if err != nil {
			return path, err
		}
		cur := val.(GraphNode)
		path = append(path, cur)
		visited[cur.key] = true
		if cur.key == goal {
			break
		}
		for _, v := range cur.adjacent {
			if in, _ := visited[v.key]; !in {
				s = s.Push(v)
			}
		}
	}
	return
}

func (n GraphNode) FindPathBFS(goal uuid.UUID) (path []GraphNode, e error) {
	visited := make(map[uuid.UUID]bool)
	var q queue.Queue
	q = q.Enqueue(n)
	for {
		if q.IsEmpty() {
			break
		}
		var val interface{}
		var err error
		q, val, err = q.Dequeue()
		if err != nil {
			return path, err
		}
		cur := val.(GraphNode)
		path = append(path, cur)
		visited[cur.key] = true
		fmt.Println("visited: ", cur.key)
		if cur.key == goal {
			break
		}
		for _, v := range cur.adjacent {
			if in, _ := visited[v.key]; !in {
				q = q.Enqueue(v)
			}
		}
	}
	return
}
