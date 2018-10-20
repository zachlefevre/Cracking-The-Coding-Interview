package stack

import (
	"errors"
)

type Stack []int

func (s Stack) Push(v int) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, int, error) {
	l := len(s)
	if l == 0 {
		return nil, 0, errors.New("Stack is empty.")
	}
	return s[:l-1], s[l-1], nil
}
