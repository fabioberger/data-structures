package queue

import (
	"errors"

	"github.com/fabioberger/data-structures/node"
)

// Queue is a struct which implements a queue using two pointers to a single linkedlist
type Queue struct {
	First *node.Node
	Last  *node.Node
}

// NewQueue creates a new queue given an initial value
func NewQueue(data int) *Queue {
	newNode := node.NewNode(data)
	q := new(Queue)
	q.First = newNode
	q.Last = newNode
	return q
}

func (q *Queue) Enqueue(data int) {
	if q.First == nil {
		newNode := node.NewNode(data)
		q.First = newNode
		q.Last = newNode
		return
	}
	newNode := node.NewNode(data)
	newNode.Next = q.First
	q.First = newNode
}

// Dequeue removes an item from the queue in the FIFO order
func (q *Queue) Dequeue() (int, error) {
	if q.First == nil {
		return 0, errors.New("No more items to dequeue")
	}
	data := q.Last.Data
	temp := q.First
	if q.First != q.Last {
		for temp.Next != q.Last {
			temp = temp.Next
		}
		q.Last = temp
		q.Last.Next = nil
	} else {
		q.First = nil
		q.Last = nil
	}
	return data, nil
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	if q.First == nil {
		return true
	}
	return false
}
