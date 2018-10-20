/*
	Given a linked list, return the k-th from the last node.

	Ideas:
		If the length of the ll is known to be n, one can easily find (n-k)th neighbor

		The ll can be converted to a slice and found that way
			easy given that I implemented Lltoas.

		recursively a function find_kth_helper(list *list, kth *node, depth int) could be called
			by find_kth.
*/
package main

import (
	"../ll"
	"fmt"
)

func findKthHelper(trav *ll.Node, kth **ll.Node, depth int, totalSize *int, k int) {
	if trav.Next == nil {
		/* The end of the list has been reached */
		*totalSize = depth
		return
	}
	findKthHelper(trav.Next, kth, depth+1, totalSize, k)
	if depth == (*totalSize - k) {
		*kth = trav
	}
}

func findKth(list ll.Sll, k int) (n *ll.Node) {
	ts := 0
	findKthHelper(list.Head, &n, 0, &ts, k)
	return
}

func main() {
	list := ll.Sll{}
	list.Append(&ll.Node{Data: 1})
	list.Append(&ll.Node{Data: 2})
	list.Append(&ll.Node{Data: 3})
	list.Append(&ll.Node{Data: 4})
	list.Append(&ll.Node{Data: 5})
	list.Append(&ll.Node{Data: 6})
	list.Append(&ll.Node{Data: 7})
	list.Append(&ll.Node{Data: 8})

	fmt.Println(findKth(list, 7))
}
