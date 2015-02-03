package graph

import (
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestFindConnectedComponents(t *testing.T) {
	g := initGraph(true)
	got := g.ConnectedComponents()
	expected := map[int][]int{1: []int{1, 2, 6, 3, 4, 5}, 2: []int{7, 8, 9, 10}}
	for i, component := range got {
		for j, vertex := range component {
			if expected[i][j] != vertex {
				t.Error("Incorrect Connected Graph Membership")
			}
		}
	}
}

func TestArticulationVectors(t *testing.T) {
	silenceOutput()
	g := initGraph(true)
	g.InitSearch()
	got := g.FindArticulationVectors(1)
	expected := []int{2, 2, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Error("Incorrect articulation vertices found")
	}
}

func TestShortestPath(t *testing.T) {
	// Test Case 1: Path Exists
	g := initGraph(true)
	start, end := 1, 5
	path, err := g.FindPath(start, end)
	if err != nil {
		t.Error("Did not find path when one exists")
	}
	expected := []int{1, 2, 3, 4, 5}
	for i, e := range expected {
		if path[i] != e {
			t.Error("Incorrect shortest path found")
		}
	}

	// Test Case 2: No Path Exists
	start, end = 6, 5
	_, err = g.FindPath(start, end)
	if err == nil {
		t.Error("Found path when none exists")
	}
}

func TestFindCycle(t *testing.T) {
	g := initGraph(true)
	g.InitSearch()
	cycleEdge := g.FindCycles(1)
	if cycleEdge[0] != 2 || cycleEdge[1] != 5 {
		t.Error("Did not find cycle that exists")
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	g := initGraph(true)
	g.InitSearch()
	got := g.BreadthFirstSearch(1)
	expected := [][]int{[]int{1}, []int{1, 2}, []int{1, 6}, []int{2}, []int{2, 3}, []int{6}, []int{3}, []int{3, 4}, []int{4}, []int{4, 5}, []int{5}, []int{5, 2}}
	for i, e := range got {
		if !reflect.DeepEqual(expected[i], e) {
			t.Error("Incorrect Breadth First Search Discovery")
		}
	}
}

func TestDepthFirstSearch(t *testing.T) {
	g := initGraph(true)
	g.InitSearch()
	got := g.DepthFirstSearch(1)
	expected := [][]int{[]int{1}, []int{1, 2}, []int{2}, []int{2, 3}, []int{3}, []int{3, 4}, []int{4}, []int{4, 5}, []int{5}, []int{5, 2}, []int{1, 6}, []int{6}}
	for i, e := range got {
		if !reflect.DeepEqual(expected[i], e) {
			t.Error("Incorrect Depth First Search Discovery")
		}
	}
}

func TestPrint(t *testing.T) {
	g := initGraph(true)
	buff := switchToBuffer()
	g.Print()
	lines := strings.Split(buff.String(), "\n")
	// This is not a comprehensive test for graph printing since the output is
	// not is a very standard format. To be improved.
	if !strings.Contains(lines[0], "Graph num edges: 9 and num vertices: 10") {
		t.Error("incorrectly read number of edges & vertices from graph")
	}
	if lines[3] != "Vert. 1 -> 2 6" {
		t.Error("Didn't parse vertices properly")
	}
}

func initGraph(directed bool) *Graph {
	g := NewGraph(directed)
	g.Read("./test_data/graph1.txt")
	return g
}

func switchToBuffer() *bytes.Buffer {
	Output = bytes.NewBuffer([]byte{})
	buff := Output.(*bytes.Buffer)
	return buff
}

func silenceOutput() {
	var err error
	if Output, err = os.Open(os.DevNull); err != nil {
		panic(err)
	}
}
