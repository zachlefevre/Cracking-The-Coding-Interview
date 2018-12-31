/*Implement an algorithm to determine if a string has all unique characters.
What if you cannot use additional data structures


  Ideas:
    The array elements could be added to a binary array (or hash table), any collisions would be a result of duplicate elements
    If no other data structures are allowed:
      The array elements could be sorted and then checked linearly.
      If no side-effects are allowed:
        The elements could be checked to be unique against every character after it. This would be O(n^2)
*/
package main

import (
	"fmt"
	"sort"
)

type s []int

func (str s) Len() int           { return len(str) }
func (str s) Less(a, b int) bool { return str[a] < str[b] }
func (str s) Swap(a, b int) {
	str[a], str[b] = str[b], str[a]
	return
}

func (str s) isUnique() bool {
	if len(str) > 128 {
		return false
	}
	sort.Sort(str)
	fmt.Println(str)
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			return false
		}
	}
	return true
}

func main() {
	str := s{1, 4, 7, 5, 3, 99, 11}
	fmt.Println(str.isUnique())
}
