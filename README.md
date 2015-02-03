[![GoDoc](http://godoc.org/github.com/fabioberger/data-structures?status.svg)](https://godoc.org/github.com/fabioberger/data-structures)

Data Structures Implemented in Go
---------------------------------

While brushing up on implementing different data structures, I decided to implement a simple version of each in Golang.

# Installation

Install like you would any other package:

```bash
go get github.com/fabioberger/data-structures
```

Add the data structure you want to use into your project's imports:

```go
import "github.com/fabioberger/data-structures/queue"
```

And off you go!

# Example Usage

## Binary Search Trees

**First import the BST package:**

```go
import "github.com/fabioberger/data-structures/bst"
```

**Create a tree and insert some values:**

```
tree := bst.NewTree(8, nil) // 2nd param is for a parent tree
tree.Insert(4)
tree.Insert(17)
tree.Insert(2)
```

**Search for a value:**

```
value, err := tree.Search(2) // binary search in O(lg n) time
if err != nil {
	fmt.Println(err) // Value not found in tree
}
fmt.Println("Found: ", value)
```

Traverse the tree ("inOrder", "preOrder" & "postOrder"):

```
traversal := []int{}
tree.Traverse("inOrder", &traversal)
fmt.Println(traversal) // [2 4 8 17]
```

Find the minimum value: 

```go
min, err := tree.Minimum()
if err != nil {
	panic(err)
}
fmt.Println(min) // 2
```
Find the maximum value:

```go
max, err := tree.Maximum()
if err != nil {
	panic(err)
}
fmt.Println(max.Item) //17
```
Delete an item:

```go
deleted, err := tree.Delete(4)
if err != nil {
	panic(err)
}
fmt.Println(deleted) // true
```

Check if the tree is balanced (i.e: no height distance > 2):

```go
isBalanced := tree.IsBalanced() 
if isBalanced {
	fmt.Println("Tree is balanced")
}
```



