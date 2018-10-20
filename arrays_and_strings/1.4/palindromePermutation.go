/* Given a string, determine if it is a permutatin of a palindrome.

Ideas:
	If the string is odd, then the only allowable number of odd-frequency characters is 1.
	If the string is even, then the only allowable number of odd-frequency characters is 0.
	Keep a frequency map to track the character and the frequency of those characters.
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
	s := "abcbakfkaaf"
	fmt.Println(s)
	fmt.Println(isPalindromePermutation(s))
}
