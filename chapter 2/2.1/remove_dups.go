/*
	Remove duplicates from an unsorted linked list.

	Ideas:
		With another data structure such as a map, one can gather the frequencies of each element on
		the linked list and remove any node which is on more than once.

		If no other data structure is allowed, the list could be sorted and traversed linearly
		If sorting is not allowed then the list would have to be traversed linearly for every node
			thus resulting in O(n^2) time
*/
package main

import "../ll"

func removeDups(list *ll.Sll) {
	s := list.Lltoas()
	stats := make(map[int]int)
	for _, n := range s {
		stats[n]++
	}

	for k := range stats {
		/*Cannot use for k,v :=... because for v is only
		evaluated each time the outter loop is run, and thus
		stats[k] and v would be out of sync and need to be managed outside
		of the inner loop*/
		for stats[k] > 1 {
			list.Remove(k)
			stats[k]--
		}
	}
}

func main() {
	list := ll.Sll{}
	list.Append(&ll.Node{Data: 5})
	list.Append(&ll.Node{Data: 6})
	list.Append(&ll.Node{Data: 5})
	list.Append(&ll.Node{Data: 7})
	list.Append(&ll.Node{Data: 8})
	list.Append(&ll.Node{Data: 1})
	list.Append(&ll.Node{Data: 1})
	list.Append(&ll.Node{Data: 8})
	list.Append(&ll.Node{Data: 8})
	list.Append(&ll.Node{Data: 8})
	list.Append(&ll.Node{Data: 8})
	list.Append(&ll.Node{Data: 8})
	list.Append(&ll.Node{Data: 8})
	list.Printll()
	removeDups(&list)
	list.Printll()

}
