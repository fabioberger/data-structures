package bst

import "testing"

func TestInOrderTraversal(t *testing.T) {
	tree := initTree()
	got := tree.Traverse("inOrder")
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 17}
	for i, val := range got {
		if val != expected[i] {
			t.Error("Traversal order incorrect")
		}
	}
}

func TestPostOrderTraversal(t *testing.T) {
	tree := initTree()
	got := tree.Traverse("postOrder")
	expected := []int{1, 3, 2, 5, 7, 6, 4, 17, 8}
	for i, val := range got {
		if val != expected[i] {
			t.Error("Traversal order incorrect")
		}
	}
}

func TestPreOrderTraversal(t *testing.T) {
	tree := initTree()
	got := tree.Traverse("preOrder")
	expected := []int{8, 4, 2, 1, 3, 6, 5, 7, 17}
	for i, val := range got {
		if val != expected[i] {
			t.Error("Traversal order incorrect")
		}
	}
}

func TestInsert(t *testing.T) {
	tree := NewTree(5, nil)
	tree.Insert(2)
	if tree.Left.Item != 2 {
		t.Error("Insert did not place the new node in the correct location")
	}
	tree.Insert(10)
	if tree.Right.Item != 10 {
		t.Error("Insert did not place the new node in the correct location")
	}
}

func TestSearch(t *testing.T) {
	tree := initTree()
	found, err := tree.Search(3)
	if err != nil {
		t.Error(err.Error())
	}
	if found.Item != 3 {
		t.Error("Did not find the correct value")
	}
}

func TestMinimum(t *testing.T) {
	tree := initTree()
	min, err := tree.Minimum()
	if err != nil {
		t.Error(err.Error())
	}
	if min.Item != 1 {
		t.Error("Did not find correct minimum")
	}
}

func TestMaximum(t *testing.T) {
	tree := initTree()
	min, err := tree.Maximum()
	if err != nil {
		t.Error(err.Error())
	}
	if min.Item != 17 {
		t.Error("Did not find correct minimum")
	}
}

func TestDelete(t *testing.T) {
	tree := initTree()
	deleted, err := tree.Delete(5)
	if err != nil {
		t.Error(err.Error())
	}
	if deleted != true {
		t.Error("Did not delete tree node")
	}
}

func TestIsBalanced(t *testing.T) {
	tree := initTree()
	if tree.IsBalanced() {
		t.Error("Says it's balanced when it isn't!")
	}
}

func initTree() *Tree {
	tree := NewTree(8, nil)
	tree.Insert(4)
	tree.Insert(17)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(5)
	tree.Insert(7)
	return tree
}
