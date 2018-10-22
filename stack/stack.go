package stack

import (
	"errors"
)

type Stack []interface{}

func (s Stack) Push(v interface{}) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, interface{}, error) {
	l := len(s)
	if l == 0 {
		return nil, 0, errors.New("Stack is empty.")
	}
	return s[:l-1], s[l-1], nil
}

func (s Stack) IsEmpty() bool {
	return len(s) <= 0
}
