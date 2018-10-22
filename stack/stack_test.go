package stack

import (
	"fmt"
	"testing"
)

func TestAdd_NewStack_StackWithNewValue(t *testing.T) {
	var s Stack
	s = s.Push(2)
	fmt.Println("s: ", s)

	s = s.Push(3)
	fmt.Println("s: ", s)

	sn := s
	s = s.Push(4)
	fmt.Println("s: ", s)
	fmt.Println("sn: ", sn)

	s = s.Push(5)
	sn = sn.Push(1)
	fmt.Println("s: ", s)
	fmt.Println("sn: ", sn)
}

func TestPop_EmptyStack_Error(t *testing.T) {
	var s Stack
	s, v, err := s.Pop()
	fmt.Println(s, v, err)
}

func TestPop_NonEmptyStack_NoErrorAndVal(t *testing.T) {
	var s Stack
	s = s.Push(3)
	s = s.Push(4)
	s, v, err := s.Pop()
	fmt.Println(s, v, err)
}
