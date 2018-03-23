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
)

func addNode(n1, n2 *ll.Node, carry int) (*ll.Node, bool) {
	var n1Quantity int
	var n2Quantity int
	if n1 != nil {
		n1Quantity = n1.Data
	}
	if n2 != nil {
		n2Quantity = n2.Data
	}
	if n1Quantity+n2Quantity+carry >= 10 {
		return &ll.Node{Data: (n1Quantity + n2Quantity + carry) - 10}, true
	}
	return &ll.Node{Data: n1Quantity + n2Quantity + carry}, false
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
		newList.Printll()
	}
	if i == nil {
		for j != nil {
			newNode, carry := addNode(i, j, carryQuantity)
			newList.Append(newNode)
			if carry {
				carryQuantity = 1
			} else {
				carryQuantity = 0
			}
			j = j.Next
			newList.Printll()
		}
	}
	if j == nil {
		for i != nil {
			newNode, carry := addNode(i, j, carryQuantity)
			newList.Append(newNode)
			if carry {
				carryQuantity = 1
			} else {
				carryQuantity = 0
			}
			i = i.Next
			newList.Printll()
		}
	}
	newNode, _ := addNode(i, j, carryQuantity)
	if newNode.Data > 0 {
		newList.Append(newNode)
	}

	return
}

func main() {
	list := ll.Sll{}
	list.Append(&ll.Node{Data: 1})
	list.Append(&ll.Node{Data: 6})
	list2 := ll.Sll{}
	list2.Append(&ll.Node{Data: 9})
	list2.Append(&ll.Node{Data: 3})
	list2.Append(&ll.Node{Data: 9})

	list.Printll()
	list2.Printll()
	sumList := addLists(list2, list)
	sumList.Printll()
}
