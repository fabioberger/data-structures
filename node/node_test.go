package node

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	ll := initLinkedList()
	got, err := convertToIntSlice(ll)
	if err != nil {
		t.Error(err.Error())
	}
	expected := []int{5, 3, 6, 8, 3, 8}
	for i, val := range got {
		if val != expected[i] {
			t.Error("Unexpected value printed")
		}
	}
}

func TestLength(t *testing.T) {
	ll := initLinkedList()
	length := ll.Length()
	if length != 6 {
		t.Error("Incorrect length found")
	}
}

func TestAppendToTail(t *testing.T) {
	ll := initLinkedList()
	ll.AppendToTail(25)
	for ll.Next != nil {
		ll = ll.Next
	}
	if ll.Data != 25 {
		t.Error("AppendToTail failed")
	}
}

func TestDeleteNode(t *testing.T) {
	ll := initLinkedList()
	ll.DeleteNode(6)
	got, err := convertToIntSlice(ll)
	if err != nil {
		t.Error(err.Error())
	}
	expected := []int{5, 3, 8, 3, 8}
	for i, v := range got {
		if v != expected[i] {
			t.Error("Node not deleted properly")
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	ll := initLinkedList()
	ll.RemoveDuplicates()
	got, err := convertToIntSlice(ll)
	if err != nil {
		t.Error(err.Error())
	}
	expected := []int{5, 3, 6, 8}
	for i, val := range got {
		if val != expected[i] {
			t.Error("Unexpected duplicate found")
		}
	}
}

func TestFindNLast(t *testing.T) {
	// Full List
	ll := initLinkedList()
	found, err := ll.FindNLast(2)
	if err != nil {
		t.Error(err.Error())
	}
	if found.Data != 3 {
		t.Error("Incorrect Nth last element found")
	}

	// Invalid input
	_, err = ll.FindNLast(-5)
	if err.Error() != "Invalid linked list or toLast parameter." {
		t.Error("Didnt catch invalid input to FindNLast")
	}

	// Empty List
	ll = new(Node)
	_, err = ll.FindNLast(2)
	if err.Error() != "Error: list size < toLast" {
		t.Error("Didnt catch empty list error for Nth Last")
	}
}

func initLinkedList() *Node {
	n := NewNode(5)
	n.AppendToTail(3)
	n.AppendToTail(6)
	n.AppendToTail(8)
	n.AppendToTail(3)
	n.AppendToTail(8)
	return n
}

func convertToIntSlice(ll *Node) ([]int, error) {
	Output = bytes.NewBuffer([]byte{})
	buff := Output.(*bytes.Buffer)
	ll.Print()
	values := strings.Split(buff.String(), "\n")
	final := []int{}
	for _, v := range values {
		if v != "" {
			intValue, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			final = append(final, intValue)
		}
	}
	return final, nil
}
