package queue

import (
	"fmt"
	"testing"
)

func TestAdd_NewQueue_QueueWithNewValue(t *testing.T) {
	var q Queue
	q = q.Enqueue(2)
	fmt.Println("q: ", q)

	q = q.Enqueue(3)
	fmt.Println("q: ", q)

	qn := q
	q = q.Enqueue(4)
	fmt.Println("q: ", q)
	fmt.Println("qn: ", qn)

	q = q.Enqueue(5)
	qn = qn.Enqueue(1)
	fmt.Println("q: ", q)
	fmt.Println("qn: ", qn)
}

func TestPop_EmptyStack_Error(t *testing.T) {
	var q Queue
	q, v, err := q.Dequeue()
	fmt.Println(q, v, err)
}

func TestPop_NonEmptyStack_NoErrorAndVal(t *testing.T) {
	var q Queue
	q = q.Enqueue(3)
	q = q.Enqueue(4)
	q, v, err := q.Dequeue()
	fmt.Println(q, v, err)
}
