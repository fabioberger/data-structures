package queue

import "testing"

func TestDequeue(t *testing.T) {
	q := initQueue()
	data, err := q.Dequeue()
	if err != nil {
		t.Error(err.Error())
	}
	if data != 1 {
		t.Error("Dequeue error")
	}
}

func TestEnqueue(t *testing.T) {
	q := new(Queue)
	q.Enqueue(5)
	data, err := q.Dequeue()
	if err != nil {
		t.Error(err.Error())
	}
	if data != 5 {
		t.Error("Enqueue error")
	}
}

func TestIsEmpty(t *testing.T) {
	q := new(Queue)
	empty := q.IsEmpty()
	if empty != true {
		t.Error("IsEmpty failed on empty queue")
	}

	q = initQueue()
	empty = q.IsEmpty()
	if empty != false {
		t.Error("IsEmpty failed on unempty queue")
	}
}

// Write queue tests here
func initQueue() *Queue {
	q := NewQueue(1)
	q.Enqueue(12)
	q.Enqueue(3)
	q.Enqueue(8)
	q.Enqueue(4)
	return q
}
