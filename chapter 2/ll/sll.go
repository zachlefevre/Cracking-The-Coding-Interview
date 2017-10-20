package ll

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type Sll struct {
	Head *Node
}

func (ll *Sll) Append(n *Node) bool {
	if ll.Head == nil {
		ll.Head = n
	} else {
		trav := ll.Head
		for ; trav.Next != nil; trav = trav.Next {
		}
		trav.Next = n
		return true
	}
	return false
}

/*
Return the node before the node containing the searched data, and the node itself
*/
func (ll *Sll) Find(d int) (*Node, *Node) {
	if ll.Head.Data == d {
		//If the Head is found with the searched data, return nil and the Head.
		return nil, ll.Head
	}
	for trav := ll.Head; trav != nil; trav = trav.Next {
		//Traverse the list, and if the node is found, return the previous node and the searched node.
		if trav.Next == nil {
			//If trav is the tail of the list, it has already been checked and is not the searched node
			//return nil nil
			return nil, nil
		}
		if trav.Next.Data == d {
			return trav, trav.Next
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
		ll.Head = ll.Head.Next
		return true
	}
	prev.Next = toRemove.Next
	return true
}

func (ll *Sll) Printll() {
	for trav := ll.Head; trav != nil; trav = trav.Next {
		fmt.Print(trav.Data, " -> ")
	}
	fmt.Println("//")
}

func (ll *Sll) Lltoas() (s []int) {
	for trav := ll.Head; trav != nil; trav = trav.Next {
		s = append(s, trav.Data)
	}
	return
}
