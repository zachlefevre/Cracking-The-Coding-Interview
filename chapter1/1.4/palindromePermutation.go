/* Given a string, determine if it is a permutatin of a palindrome.

Ideas:
  If the string has less than, or more than 1 character with an odd frequency, it is not a palindrome permutation.
  Sort the string, either in-place or in a new array.
  If the sorted array has more than one character with an odd frequency, it is not a palindrome permutation.
  If the sorted array has less than one character with an odd frequency, it is not a palindrome permutation.

*/
package main

import (
	"fmt"
)

func isPalindromePermutation(s string) bool {
	hmap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hmap[s[i]]++
	}

	var oddCount int
	for _, value := range hmap {
		if value%2 != 0 {
			oddCount++
		}
	}
	var allowableCount int
	if len(s)%2 == 0 {
		allowableCount = 0
	} else {
		allowableCount = 1
	}
	if oddCount != allowableCount {
		return false
	} else {
		return true
	}
}
func main() {
	s := "abcba"
	fmt.Println(s)
	fmt.Println(isPalindromePermutation(s))
}
