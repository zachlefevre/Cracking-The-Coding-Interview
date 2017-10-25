/*
	Given two linked lists whose individual nodes represent single digits, and whose
		entirety represent a multi-digit number. Sum two linked lists together
		and return their sum as a linked list in the same format
			The digits are in reverse order, so the head is the
			least significant digit

	Ideas:
		This can be done iteratively and recursively
		Recursively:
			The algorithm would walk the list adding digits. Create a node from the least significant
			digit, pass the most significant digit into a "carry digit" parameter, and
			call the recursive algorithm on the next node, adding the carry digit.
			The base case (one reaching nil), would be defined such that the the carry digit is added
			to whatever is left of the other LL, and when both reach nil the carry digit is added to 0.
		Iteratively:
			A loop over the two LL nodes would add digits, a "carry digit" variable would be defined
			and added to the variables in each iteration.

			The lists could be converted to an array which would make calculating the len of the lists
			easier, which would in turn make the logic simpler (because you would not have to manage the
			first list ending before the second list)
*/

package main

import (
	"../ll"
	"fmt"
)

func addNode(n1, n2 *ll.Node, carry int) (*ll.Node, bool) {

	if n1.Data+n2.Data+carry >= 10 {
		return &ll.Node{Data: (n2.Data + n1.Data + carry) - 10}, true
	}
	return &ll.Node{Data: n2.Data + n1.Data + carry}, false
}

func addLists(list1, list2 ll.Sll) (newList ll.Sll) {
	i := list1.Head
	j := list2.Head
	var carryQuantity int
	for i != nil && j != nil {

		newNode, carry := addNode(i, j, carryQuantity)
		newList.Append(newNode)
		if carry {
			carryQuantity = 1
		} else {
			carryQuantity = 0
		}
		i = i.Next
		j = j.Next
		if i == nil {
			for j != nil {
				//todo (zlefevre): Bug whereby 7.next becomes 7
				newList.Append(j)
				fmt.Println("j: ", j, "\t\tj.next: ", j.Next)
				j = j.Next
			}
		}
	}
	newList.Printll()
	return
}

func main() {
	list := ll.Sll{}
	list.Append(&ll.Node{Data: 4})
	list.Append(&ll.Node{Data: 6})
	list.Append(&ll.Node{Data: 5})
	list.Append(&ll.Node{Data: 1})
	list.Printll()
	list2 := ll.Sll{}
	list2.Append(&ll.Node{Data: 4})
	list2.Append(&ll.Node{Data: 6})
	list2.Append(&ll.Node{Data: 5})
	list2.Append(&ll.Node{Data: 2})
	list2.Append(&ll.Node{Data: 9})
	list2.Append(&ll.Node{Data: 7})

	addLists(list, list2)
}
