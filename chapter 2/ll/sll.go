package ll

import (
	"fmt"
)

type Node struct {
	Data int
	next *Node
}

type Sll struct {
	head *Node
}

func (ll *Sll) Append(n *Node) bool {
	if ll.head == nil {
		ll.head = n
	} else {
		trav := ll.head
		for ; trav.next != nil; trav = trav.next {
		}
		trav.next = n
		return true
	}
	return false
}

func (ll *Sll) Printll() {
	for trav := ll.head; trav != nil; trav = trav.next {
		fmt.Print(trav.Data, " -> ")
	}
	fmt.Println("//")
}
