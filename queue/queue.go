package queue

import (
	"errors"
)

type Queue []interface{}

func (q Queue) Enqueue(v interface{}) Queue {
	return append(q, v)
}

func (q Queue) Dequeue() (Queue, interface{}, error) {
	l := len(q)
	if l == 0 {
		return nil, nil, errors.New("Queue is empty")
	}
	return q[1:], q[0], nil
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
