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

/*
Return the node before the node containing the searched data, and the node itself
*/
func (ll *Sll) Find(d int) (*Node, *Node) {
	if ll.head.Data == d {
		//If the head is found with the searched data, return nil and the head.
		return nil, ll.head
	}
	for trav := ll.head; trav != nil; trav = trav.next {
		//Traverse the list, and if the node is found, return the previous node and the searched node.
		if trav.next == nil {
			//If trav is the tail of the list, it has already been checked and is not the searched node
			//return nil nil
			return nil, nil
		}
		if trav.next.Data == d {
			return trav, trav.next
		}
	}
	return nil, nil
}

func (ll *Sll) Remove(d int) bool {
	prev, toRemove := ll.Find(d)
	if prev == nil && toRemove == nil {
		return false
	}
	if prev == nil {
		ll.head = ll.head.next
		return true
	}
	prev.next = toRemove.next
	return true
}

func (ll *Sll) Printll() {
	for trav := ll.head; trav != nil; trav = trav.next {
		fmt.Print(trav.Data, " -> ")
	}
	fmt.Println("//")
}

func (ll *Sll) Lltoas() (s []int) {
	for trav := ll.head; trav != nil; trav = trav.next {
		s = append(s, trav.Data)
	}
	return
}
