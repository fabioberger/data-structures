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

**Traverse the tree ("inOrder", "preOrder" & "postOrder"):**

```
traversal := []int{}
tree.Traverse("inOrder", &traversal)
fmt.Println(traversal) // [2 4 8 17]
```

**Find the minimum value:** 

```go
min, err := tree.Minimum()
if err != nil {
	panic(err)
}
fmt.Println(min) // 2
```
**Find the maximum value:**

```go
max, err := tree.Maximum()
if err != nil {
	panic(err)
}
fmt.Println(max.Item) //17
```
**Delete an item:**

```go
deleted, err := tree.Delete(4)
if err != nil {
	panic(err)
}
fmt.Println(deleted) // true
```

**Check if the tree is balanced (i.e: no height distance > 2):**

```go
isBalanced := tree.IsBalanced() 
if isBalanced {
	fmt.Println("Tree is balanced")
}
```

## Stacks

**Import the package:**

```go
import "github.com/fabioberger/data-structures/stack"
```

**Create a new stack and add items to it:**

```go
s := NewStack()
s.Push(1)
s.Push(5)
s.Push(7)
s.Push(10)
```

**Peek at the top value in the stack:**

```go
peeked := s.Peek()
fmt.Println(peeked) // 10
```

**Pop an item off the top of the stack:**

```go
item, err := s.Pop()
if err != nil {
	fmt.Println(err) // No more items to pop!
}
fmt.Println(item) // 10
```

**Check if the stack is empty:**

```go
empty := s.IsEmpty()
if empty {
	fmt.Println("Stack is empty")
}
```

## Queues

**Import the package:**

```go
import "github.com/fabioberger/data-structures/queue"
```

**Create a new queue and add items to it:**

```go
q := NewQueue(1)
q.Enqueue(12)
q.Enqueue(3)
q.Enqueue(8)
q.Enqueue(4)
```

**Dequeue an item off the end of the queue:**

```go
data, err := q.Dequeue()
if err != nil {
	fmt.Println(err) // No more items to dequeue
}
fmt.Println(data) // 1
```

**Check if the queue is empty:**

```go
empty := q.IsEmpty()
fmt.Println(empty) // false
```

## Singly Linked Lists

**Import the package:**

```go
import "github.com/fabioberger/data-structures/node"
```

**Create a new linked list and add items to it:**

```go
ll := NewNode(5)
ll.Append(3)
ll.Append(6)
ll.Append(8)
ll.Append(3)
ll.Append(8)
```

**Calculate length of the linked list:**

```go
length := ll.Length()
fmt.Println(length) //6
```

**Delete a node by value (removes first node found with this value):**

```go
ll.DeleteNode(6)
```

**Remove all duplicates from the linked list:**

```go
ll.RemoveDuplicates()
```

**Find nth to last node:**

```go
found, err := ll.FindNLast(2)
if err != nil {
	panic(err)
}
fmt.Println(found) // 3
```


