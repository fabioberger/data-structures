package stack

import (
	"errors"

	"github.com/fabioberger/data-structures/node"
)

// The stack struct implements a stack using a singly-linkedlist
type Stack struct {
	Top *node.Node
}

// NewStack initializes a new empty stack
func NewStack() *Stack {
	s := new(Stack)
	s.Top = nil
	return s
}

// ISEmpty checks if the given stack is empty
func (s *Stack) IsEmpty() bool {
	if s.Top == nil {
		return true
	}
	return false
}

// Peek returns the top item in the stack without removing it
func (s *Stack) Peek() int {
	if s.Top != nil {
		return s.Top.Data
	}
	return -1
}

// Push adds a new data item to the top of the stack
func (s *Stack) Push(data int) {
	n := node.NewNode(data)
	n.Next = s.Top
	s.Top = n
}

// PushNode pushes a node onto the top of the stack
func (s *Stack) PushNode(n *node.Node) {
	n.Next = s.Top
	s.Top = n
}

// PopNode removes and returns the top node from the stack
func (s *Stack) PopNode() (*node.Node, error) {
	if s.Top == nil {
		return nil, errors.New("No more items to pop!")
	}
	node := s.Top
	s.Top = s.Top.Next
	return node, nil
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() (int, error) {
	if s.Top == nil {
		return 0, errors.New("No more items to pop!")
	}
	data := s.Top.Data
	s.Top = s.Top.Next
	return data, nil
}
