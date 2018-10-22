package graph

import "github.com/zachlefevre/Cracking-The-Coding-Interview/stack"

type GraphNode struct {
	key      int
	value    interface{}
	adjacent []GraphNode
}

func (n *GraphNode) AddNeighbor(a GraphNode) {
	n.adjacent = append(n.adjacent, a)
}

func (n GraphNode) FindPathDFS(goal int) (path []GraphNode, e error) {
	visited := make(map[int]bool)
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
