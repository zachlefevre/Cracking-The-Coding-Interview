/*
	Given two strings, check to see if they are one  "edit" away from one another, where an "edit" is
		defined as a removal of a character, insertion of a character, and replacing of a character

	Ideas
		Brute Force:
			Find all permutations of the string which fulfills a single edit as described above.
				Compare the edited string against these permutations

		Intelligent:
			Check the lengths of the string to determine if an insertion or deletion is the edit.
				If the lengths match, return oneChangeHelper()
				Note: An insert onto A to create B is the same as removal onto B to create A,
					therefore oneInsertHelper(A,B) =  oneRemovalHelper(B, A)
				If the length of string B is longer than string A, return oneRemovalHelper(A,B)
				If the length of string B is shorter than string A, return oneRemovalHelper(B,A)

		Not polished
*/

package main

import (
	"fmt"
)

func oneChangeHelper(a, b []byte) bool {
	edit := 0
	for i := range a {
		if a[i] != b[i] {
			edit++
		}
	}
	if edit == 1 {
		return true
	}
	return false
}

func oneRemovalHelper(a, b []byte) bool {
	/*len(b) < len(a)*/
	removed := 0
	j := 0
	i := 0
	for i < len(a) {
		if a[i] != b[j] {
			removed++
			i++
		} else {
			i++
			j++
		}
		if j >= len(b) {
			/* If we reach the end of j, and the difference in length between the two strings is 1,
			then the only possibility is that the last character is missing*/
			return true
		}
	}
	if removed == 1 {
		return true
	}
	return false
}

func main() {
	str1 := "zachary"
	str2 := "Zachary"
	if len(str1) == len(str2) {
		fmt.Println(oneChangeHelper([]byte(str1), []byte(str2)))
	} else if len(str1) == (len(str2) + 1) {
		fmt.Println(oneRemovalHelper([]byte(str1), []byte(str2)))
	} else if len(str2) == (len(str1) + 1) {
		fmt.Println(oneRemovalHelper([]byte(str2), []byte(str1)))
	} else {
		fmt.Println("Difference too great to be one edit away")
	}
}
