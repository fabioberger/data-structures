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

## Graphs

**Warning: Still a work in progress**

**Import the package:**

```go
import "github.com/fabioberger/data-structures/graph"
```

**Create a new graph and add edges via a file:**

```go
g := NewGraph(true) // true for a directed graph
g.Read("./test_data/graph1.txt")
```
where the file contains two ints per line representing the two vertices of an edge

**Print the graph:**

```go
g.Print()
/*
Expected Output:
Graph num edges: 9 and num vertices: 10 
Directed? true
Adjacency List:
Vert. 1 -> 2 6
Vert. 2 -> 3
Vert. 3 -> 4
Vert. 4 -> 5
Vert. 5 -> 2
Vert. 6 ->
Vert. 7 -> 8
Vert. 8 -> 9
Vert. 9 -> 10
Vert. 10 ->
*/
```

**Breadth first search:**

```go
g.InitSearch()
traversal := g.BreadthFirstSearch(1) // start at vertice 1
fmt.Println(traversal) // [[1] [1 2] [1 6] [2] [2 3] [6] [3] [3 4] [4] [4 5] [5] [5 2]]
```
Single values are discovered vertices, double values are discovered edges

**Find the shortest past:**

```go
path, err := g.FindPath(start, end)
if err != nil {
	fmt.Println(err) // No Path exists
}
fmt.Println(path) // [1 2 3 4 5]
```

**Find all connected components of the graph:**

```go
components := g.ConnectedComponents()
fmt.Println(components) // map[1:[1 2 6 3 4 5] 2:[7 8 9 10]] (two separate, connected components)
```

**Depth first search:**

```go
g.InitSearch()
got := g.DepthFirstSearch(1) // [[1] [1 2] [2] [2 3] [3] [3 4] [4] [4 5] [5] [5 2] [1 6] [6]]
```
Single values are discovered vertices, double values are discovered edges

**Find any cycles in the graph:**

```go
g.InitSearch()
cycleEdge, err := g.FindCycles(1)
if err != nil {
	panic(err) // Did not find the existing cycle
}
fmt.Println(cycleEdge) // [2 5]
```

**Find all articulation vectors:**

```go
g.InitSearch()
articulationVectors := g.FindArticulationVectors(1)
fmt.Println(articulationVectors) // [2 2 1]
```
