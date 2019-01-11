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
