package node

import (
	"errors"
	"fmt"
)

// Node is a struct that models a singly-linkedlist
type Node struct {
	Next *Node
	Data int
}

// NewNode creates a new linkedlist given an initial value
func NewNode(data int) *Node {
	return &Node{
		Next: nil,
		Data: data,
	}
}

// Print traverses the linkedlist and prints out each node
func (n *Node) Print() {
	for n != nil {
		fmt.Println(n)
		n = n.Next
	}
}

// Length returns the number of nodes in the linkedlist
func (n *Node) Length() int {
	len := 0
	for n != nil {
		len++
		n = n.Next
	}
	return len
}

// AppendToTail adds a node to the end of a linkedlist
func (n *Node) AppendToTail(data int) {
	end := NewNode(data)
	for n.Next != nil {
		n = n.Next
	}
	n.Next = end
}

// DeleteNode remove the first node with the supplied value from the linkedlist
func (head *Node) DeleteNode(data int) {
	node := head
	if node.Data == data {
		*head = *head.Next
		return
	}
	for node.Next != nil {
		if node.Next.Data == data {
			node.Next = node.Next.Next
			return
		}
		node = node.Next
	}
}

// RemoveDuplicates removes any nodes with same value from linkedlist
func (n *Node) RemoveDuplicates() {
	if n == nil {
		return
	}
	prev := n
	curr := n.Next
	for curr != nil {
		runner := n
		for runner != curr {
			if runner.Data == curr.Data {
				tmp := curr.Next
				prev.Next = tmp
				curr = tmp
				break
			}
			runner = runner.Next
		}
		if runner == curr {
			prev = curr
			curr = curr.Next
		}
	}
}

// Find node that is 'toLast' nodes from the end of the linkedlist
func (n *Node) FindNLast(toLast int) (*Node, error) {
	if n == nil || toLast < 0 {
		return nil, errors.New("Invalid linked list or toLast parameter.")
	}
	follower := n
	leader := n
	for i := 0; i < toLast-1; i++ {
		if leader == nil {
			return nil, errors.New("Error: list size < toLast")
		}
		leader = leader.Next
	}
	for leader.Next != nil {
		follower = follower.Next
		leader = leader.Next
	}
	return follower, nil
}
