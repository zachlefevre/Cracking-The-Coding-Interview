package main

import "fmt"

type nodeData struct {
	data int
}

func (n *nodeData) less(o *nodeData) bool { return n.data < o.data }
func (n *nodeData) println()              { fmt.Println(n.data) }

type treeNode struct {
	data  *nodeData
	left  *treeNode
	right *treeNode
}

func (t *treeNode) less(o *treeNode) bool { return t.data.less(o.data) }
func (t *treeNode) addNode(n *nodeData) bool {
	c := &treeNode{data: n}

	if !t.less(c) {
		if t.left == nil {
			t.left = c
		} else {
			return t.left.addNode(n)
		}
	} else {
		if t.right == nil {
			t.right = c
		} else {
			return t.right.addNode(n)
		}
	}
	return true
}

func (t *treeNode) inOrder() {
	if t == nil {
		return
	}
	t.left.inOrder()
	t.data.println()
	t.right.inOrder()
}

func (t *treeNode) preOrder() {
	if t == nil {
		return
	}
	t.data.println()
	t.left.preOrder()
	t.right.preOrder()
}

func (t *treeNode) postOrder() {
	if t == nil {
		return
	}
	t.left.postOrder()
	t.right.postOrder()
	t.data.println()
}

func main() {
	root := &treeNode{data: &nodeData{10}}
	root.addNode(&nodeData{15})
	root.addNode(&nodeData{13})
	root.addNode(&nodeData{11})

	root.inOrder()

}
