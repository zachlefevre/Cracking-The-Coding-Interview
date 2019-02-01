package isUnique

import (
	"errors"
	"sort"
)

type s []int

func (str s) Len() int           { return len(str) }
func (str s) Less(a, b int) bool { return str[a] < str[b] }
func (str s) Swap(a, b int) {
	str[a], str[b] = str[b], str[a]
	return

}

func (str s) isUnique() (bool, error) {
	if len(str) > 128 {
		return false, nil
	}
	if len(str) <= 0 {
		return false, errors.New("string cannot be empty")
	}
	sort.Sort(str)
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			return false, nil
		}
	}
	return true, nil
}
