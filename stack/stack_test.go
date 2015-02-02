package stack

import (
	"testing"
)

func TestPop(t *testing.T) {
	s := initStack()
	item, err := s.Pop()
	if err != nil {
		t.Error(err.Error())
	}
	if item != 10 {
		t.Error("Stack pop failed")
	}
}

func TestPush(t *testing.T) {
	s := initStack()
	s.Push(25)
	peeked := s.Peek()
	if peeked != 25 {
		t.Error("stack did not properly push value")
	}
}

func TestIsEmpty(t *testing.T) {
	s := new(Stack)
	empty := s.IsEmpty()
	if empty != true {
		t.Error("Stack IsEmpty failed on empty")
	}

	s = initStack()
	empty = s.IsEmpty()
	if empty != false {
		t.Error("Stack IsEmpty failed on full")
	}
}

func TestPeek(t *testing.T) {
	s := initStack()
	peeked := s.Peek()
	if peeked != 10 {
		t.Error("Peek did not return correct value")
	}
	value, err := s.Pop()
	if err != nil {
		t.Error(err.Error())
	}
	if value != 10 {
		t.Error("Peek modified the stack...")
	}
}

func initStack() *Stack {
	s := NewStack()
	s.Push(1)
	s.Push(5)
	s.Push(7)
	s.Push(10)
	return s
}
