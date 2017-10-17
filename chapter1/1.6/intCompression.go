/*
The string compression should, given a string, return a string containing a character followed by the
	repetition of that character iff the new string is shorter than the original string.
	aaaaabbc becomes a5b2c1
	aaabbc becomes a3b2c1, but aaabbc is returned because the return value is not less than the original string.

Ideas:
	A map could keep track of the frequency of each character, the new string could be constructed, and the new
		string could be returned iff the len is less than the old string.
			Problems: This will not work because the same characters are not necessarily sequential

	Make a slice of slices which contain the intermediate information regarding each character, and it's immediate
		frequency. Compress(s) should get the stats of s with this function, and then flatten it into a new string.

	Keep a running count of the local frequency of the current character. When a new character is found
		append the character and it's count, then set the frequency counter to 0.

Below I use ints for simplicity. A string in golang is a slice of bytes, and this implementation is analagous.
*/

package main

import (
	"fmt"
)

func arrStats(s []int) [][]int {
	nArr := [][]int{[]int{s[0], 1}}
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			nArr = append(nArr, []int{s[i], 1})
		} else {
			nArr[len(nArr)-1][1]++
		}
	}
	return nArr
}

func compress(s []int) []int {
	stats := arrStats(s)
	nString := make([]int, 0)
	for _, i := range stats {
		nString = append(nString, i[0])
		nString = append(nString, i[1])
	}

	if len(s) <= len(nString) {
		return s
	}
	return nString
}

func main() {
	s := []int{1, 1, 4, 4, 4, 4, 4, 4, 4, 4, 4, 7, 4, 8, 8}
	fmt.Println(s)
	fmt.Println(compress(s))
}
