// Binary Search Tree Implementation
package bst

import (
	"errors"
	"io"
	"math"
	"os"
)

var Output io.Writer = os.Stdout

// Tree struct defines a tree node that store an int Item
// It is a recursive data structure with tree nodes as children and parent
type Tree struct {
	Item   int
	Left   *Tree
	Right  *Tree
	Parent *Tree
}

// NewTree constructs a new tree node
func NewTree(item int, parent *Tree) *Tree {
	t := new(Tree)
	t.Item = item
	t.Parent = parent
	return t
}

// Insert adds a node to the BST in its correct location
func (t *Tree) Insert(item int) {
	var nt *Tree
	var side string
	if item < t.Item {
		nt = t.Left
		side = "left"
	} else {
		nt = t.Right
		side = "right"
	}

	if nt == nil {
		temp := NewTree(item, t)
		if side == "left" {
			t.Left = temp
		} else {
			t.Right = temp
		}
		return
	}

	nt.Insert(item)
}

// Search will find an item within a BST in O(log n) time
func (t *Tree) Search(item int) (*Tree, error) {
	if t == nil {
		return nil, errors.New("Reached nil node!")
	}
	if t.Item == item {
		return t, nil
	}

	if item < t.Item {
		found, err := t.Left.Search(item)
		if err == nil {
			return found, nil
		}
	} else {
		found, err := t.Right.Search(item)
		if err == nil {
			return found, nil
		}
	}
	return nil, errors.New("Node not found")
}

// Minimum finds the minimum value in a given tree
// The min item in a BST will be the left most node
func (t *Tree) Minimum() (*Tree, error) {
	if t == nil {
		return nil, errors.New("Tree is Empty!")
	}

	min := t
	for min.Left != nil {
		min = min.Left
	}
	return min, nil
}

// Maximum finds the maximum value in a given tree
// The max item in a BST will be the right most node
func (t *Tree) Maximum() (*Tree, error) {
	if t == nil {
		return nil, errors.New("Tree is Empty!")
	}

	max := t
	for max.Right != nil {
		max = max.Right
	}
	return max, nil
}

// Traverse will visit each node in a BST, printing each one as it goes
// This method requires specifying the kind of traversal to perform:
// either inOrder, preOrder or postOrder traversal
func (t *Tree) Traverse(kind string) []int {
	collector := []int{}
	t.TraverseRecursively(kind, &collector)
	return collector
}

func (t *Tree) TraverseRecursively(kind string, collector *[]int) {
	if t != nil {
		switch {
		case kind == "inOrder":
			t.Left.TraverseRecursively("inOrder", collector)
			(*collector) = append((*collector), t.Item)
			t.Right.TraverseRecursively("inOrder", collector)
		case kind == "preOrder":
			(*collector) = append((*collector), t.Item)
			t.Left.TraverseRecursively("preOrder", collector)
			t.Right.TraverseRecursively("preOrder", collector)
		case kind == "postOrder":
			t.Left.TraverseRecursively("postOrder", collector)
			t.Right.TraverseRecursively("postOrder", collector)
			(*collector) = append((*collector), t.Item)
		}
	}
}

// Delete will remote a node from the BST
// It covers all the edge cases: nodes with no decendents, two children and
// only one child
func (t *Tree) Delete(item int) (bool, error) {
	node, err := t.Search(item)
	if err != nil {
		return false, err
	}
	switch {
	// Case 1: Node has no descendents
	case node.Left == nil && node.Right == nil:
		if node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
	// Case 2: Node has two children nodes
	case node.Left != nil && node.Right != nil:
		successor, err := node.Right.Minimum()
		if err != nil {
			return false, err
		}
		node.Item = successor.Item
		successor.Delete(successor.Item)
	// Case 3: Node has only one child
	default:
		if node.Left != nil {
			node.Parent.Left = node.Left
			temp := node.Left
			temp.Parent = node.Parent
			node = temp
		} else {
			node.Parent.Right = node.Right
			temp := node.Right
			temp.Parent = node.Parent
			node = temp
		}
	}
	return true, nil
}

// IsBalanced checks if the BST is balanced (i.e: no height distance > 2)
func (t *Tree) IsBalanced() bool {
	max := t.MaxDepth()
	min := t.MinDepth()
	if (max - min) < 2 {
		return true
	}
	return false
}

// MaxDepth returns the height of the furthest leaf node
func (t *Tree) MaxDepth() float64 {
	if t == nil {
		return 0.0
	}
	return 1.0 + math.Max(t.Left.MaxDepth(), t.Right.MaxDepth())
}

// MinDepth returns the height of the shortest leaf node
func (t *Tree) MinDepth() float64 {
	if t == nil {
		return 0.0
	}
	return 1.0 + math.Min(t.Left.MaxDepth(), t.Right.MaxDepth())
}
