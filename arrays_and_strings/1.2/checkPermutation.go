package main


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

func isPermutation_Sort(s1, s2 s) bool {
	if len(s1) != len(s2) {
		return false
	}
	sort.Sort(s1)
	sort.Sort(s2)
	if reflect.DeepEqual(s1, s2) {
		return true
	}
	return false
}

func isPermutation_Map(s1, s2 s) bool {
	if len(s1) != len(s2) {
		return false
	}

	lookup := make(map[int]int)

	for _, el := range s1 {
		lookup[el] = lookup[el] + 1
	}
	for _, el := range s2 {
		lookup[el] = lookup[el] - 1
		if lookup[el] < 0 {
			return false
		}
	}
	return true
}

func main() {
	str1 := s{1, 7, 53, 6, 8, 88, 9, 64, 5, 1}
	str2 := s{1, 7, 53, 88, 6, 8, 9, 64, 5, 2}
	fmt.Println("first string: ", str1)
	fmt.Println("second string: ", str2)
	if isPermutation_Map(str1, str2) {
		fmt.Println("permutation")
	} else {
		fmt.Println("no permutation")
	}
}
