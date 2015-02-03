package main

import (
	"fmt"

	"github.com/fabioberger/data-structures/bst"
)

func main() {
	// Create the tree
	tree := bst.NewTree(8, nil)

	// Insert values
	tree.Insert(4)
	tree.Insert(17)
	tree.Insert(2)

	// Search for a value
	value, err := tree.Search(2)
	if err != nil {
		panic(err) // Value not in tree
	}
	fmt.Println("Found: ", value)

	// inOrder Traversal (works with preOrder & postOrder as well)
	traversal := []int{}
	tree.Traverse("inOrder", &traversal)
	fmt.Println(traversal) // [2 4 8 17]

	min, err := tree.Minimum()
	if err != nil {
		panic(err)
	}
	fmt.Println(min.Item) // 2

	max, err := tree.Maximum()
	if err != nil {
		panic(err)
	}
	fmt.Println(max.Item) //

	deleted, err := tree.Delete(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(deleted) // true

	isBalanced := tree.IsBalanced()
	if isBalanced {
		fmt.Println("Tree is balanced")
	}
}
