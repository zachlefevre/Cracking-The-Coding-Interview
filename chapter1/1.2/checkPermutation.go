package main

/*
Given two strings, write a method to decide if one is a permutation of the other

Ideas:
  If modification is allowed:
    The second second string can be ran though upon each element of the first string, and the first matching character removed
      If the string is not "" when the end of the first string is reached then it is not a permutation
      If the string is "" before the end of the first string is reached, then it is not a permutation
    The two strings could be sorted (They could also be sorted out of place to protect from modification),
      and if the sorted strings are equivalent to one another, the second string was a permutation of the first.

  If other data structures are allowed:
    An int array could be created which the codomain of a map from characters to their frequency:
      The first string would be indexed and each character incrementing a unique element on the int array.
      The second string would be indexed and each character decrementing a unique element on the int array
      If every element of the int array is 0, then the second string is a permutation of the first.
      Otherwise it is not
    A hash table could be created:
      The same process as above could be performed


*/
import (
	"fmt"
	"reflect"
	"sort"
)

type s []int

func (str s) Len() int           { return len(str) }
func (str s) Less(a, b int) bool { return str[a] < str[b] }
func (str s) Swap(a, b int) {
	str[a], str[b] = str[b], str[a]
	return
}

func main() {
	str1 := s{1, 7, 53, 6, 8, 88, 9, 64, 5}
	str2 := s{1, 7, 53, 88, 6, 8, 9, 64, 5}
	fmt.Println("first string: ", str1)
	fmt.Println("second string: ", str2)
	sort.Sort(str1)
	sort.Sort(str2)
	if reflect.DeepEqual(str1, str2) {
		fmt.Println("Permutations of one another")
	} else {
		fmt.Println("Not permutations of one another")
	}
}
