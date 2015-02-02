package graph

import (
	"fmt"
	"io"
	"os"

	"github.com/fabioberger/data-structures/queue"
)

// Enum of Vertice States (i.e DISCOVERED, UNDISCOVERED, PROCESSED)
type VerticeState int

const (
	UNDISCOVERED = 1 + iota
	DISCOVERED
	PROCESSED
)

// Enum of edge classifications (i.e TREE, BACK, FORWARD, CROSS EDGE)
type EdgeType int

const (
	TREE = 1 + iota
	BACK
	FORWARD
	CROSS
)

// GraphProcessor is an interface that must be satisfied by any specific DFS or BFS
// Processing task.
type GraphProcessor interface {
	processEdge(*Graph, int, int)
	processVertexEarly(*Graph, int)
	processVertexLate(*Graph, int)
}

// An EdgeNode represents a singe vertice of a graphs adjacency list
type EdgeNode struct {
	Y      int
	Weight int
	Next   *EdgeNode
}

// A Graph contains all the data structures necessary to describe the properties of a Graph
type Graph struct {
	Edges             map[int]*EdgeNode    //Adjacency list of edges
	Degree            map[int]int          // degree of each edge
	nVertices         int                  // Number of vertices
	nEdges            int                  // Number of Edges
	Directed          bool                 // Is the graph directed or undirected?
	State             map[int]VerticeState // State of each vertice (discovered, processed, etc.)
	Parent            map[int]int          // Who is the parent of a given vertice
	Time              int                  // Time keeper during graph traversals
	EntryTime         map[int]int          // Time when vertices were first entered
	ExitTime          map[int]int          // Time when vertices were exited
	ReachableAncestor map[int]int          // Earliest reachable ancestor of a vertice
	TreeOutDegree     map[int]int
	Finished          bool // Graph traversal end condition reached
}

// NewGraph instantiates a new Graph struct with sensible default values
func NewGraph(directed bool) *Graph {
	g := new(Graph)
	g.nVertices = 0
	g.nEdges = 0
	g.Directed = directed
	g.Finished = false
	g.Edges = make(map[int]*EdgeNode)
	g.Degree = make(map[int]int)
	return g
}

// InsertEdge adds a node to the adjacency list of the Graph and updates
// the necessary edge and degree counts
// x is adjacent edge to y which is the Id of the new edge being inserted
func (g *Graph) InsertEdge(x, y int, directed bool) {
	p := new(EdgeNode)
	p.Weight = 0
	p.Y = y // value of the new adjacent vertex to x
	p.Next = g.Edges[x]

	g.Edges[x] = p
	g.Degree[x]++

	if directed == false {
		g.InsertEdge(y, x, true)
	} else {
		g.nEdges++
	}
}

// Print outputs a representation of the Graph based on its adjacency list
func (g *Graph) Print() {
	fmt.Printf("Graph num edges: %v and num vertices: %v \n", g.nEdges, g.nVertices)
	fmt.Println("Directed? ", g.Directed)
	fmt.Println("Adjacency List:")
	var temp *EdgeNode
	for i := 1; i <= g.nVertices; i++ {
		fmt.Printf("Vert. %v ->", i)
		temp = g.Edges[i]
		for temp != nil {
			fmt.Printf(" %v", temp.Y)
			temp = temp.Next
		}
		fmt.Println("")
	}
}

// Read in values from a file to construct a graph
// The file includes one edge per line described as two ints, the two vertices
// that make up an edge
func (g *Graph) Read(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	seenVertices := make(map[int]bool)

	var x, y int
	for {
		_, err := fmt.Fscanf(file, "%v %v", &x, &y)
		if err == io.EOF {
			break
		}
		g.InsertEdge(x, y, g.Directed)
		if _, ok := seenVertices[x]; !ok {
			seenVertices[x] = true
			g.nVertices++
		}
		if _, ok := seenVertices[y]; !ok {
			seenVertices[y] = true
			g.nVertices++
		}
	}
}

// InitSearch re-initializes the State and Parent relationships of the graph vertices
func (g *Graph) InitSearch() {
	g.EntryTime = make(map[int]int)
	g.ExitTime = make(map[int]int)
	g.ReachableAncestor = make(map[int]int)
	g.TreeOutDegree = make(map[int]int)
	g.State = make(map[int]VerticeState)
	g.Parent = make(map[int]int)
	for i := 1; i <= g.nVertices; i++ {
		g.State[i] = UNDISCOVERED
		g.Parent[i] = -1
	}
	g.Time = 0
	g.Finished = false
}

// bfs is a Breadth-first search through a graph while allowing the caller to
// define how to process each iteration of the traversal by passing in a struct
// that implements GraphProcessor
// Runs in linear O(n+m) time
// Used a queue (FIFO) as node discovery order
func (g *Graph) bfs(start int, t GraphProcessor) {
	q := queue.NewQueue(start)
	g.State[start] = DISCOVERED

	var v, y int
	var edgeNode *EdgeNode
	for !q.IsEmpty() {
		v, _ = q.Dequeue() //shouldnt hit an error here b/c of surrounding for loop
		t.processVertexEarly(g, v)
		g.State[v] = PROCESSED
		edgeNode = g.Edges[v]
		for edgeNode != nil { // Now process all of the vertex's adjacent vertices
			y = edgeNode.Y
			// If edge not processed yet or its a directed graph (now exploring the edge
			// in the correct direction) then process it
			if g.State[y] != PROCESSED || g.Directed {
				t.processEdge(g, v, y)
			}
			if g.State[y] == UNDISCOVERED {
				q.Enqueue(y)
				g.State[y] = DISCOVERED
				g.Parent[y] = v
			}
			edgeNode = edgeNode.Next
		}
		t.processVertexLate(g, v)
	}
}

// QuietTraversal implements the GraphProcessor interface without outputing anything
// during the traversal. It still sets the correct State & Parent relationships for the vertices
type QuietTraversal struct {
}

func (t *QuietTraversal) processVertexEarly(g *Graph, v int) {
	// Do nothing here
}

func (t *QuietTraversal) processVertexLate(g *Graph, v int) {
	// Do nothing here
}

func (t *QuietTraversal) processEdge(g *Graph, x int, y int) {
	// Do nothing here
}

// FindPath finds the shortest path between start and end in an unweighted graph
// This only works if BFS first performed with start as its start
func (g *Graph) FindPath(start, end int) {
	t := new(QuietTraversal)
	g.InitSearch()
	g.bfs(start, t)
	g.traversePath(start, end)
}

// Traverse the shortest path between two nodes recursively printing the path
func (g *Graph) traversePath(start, end int) {
	if g.Parent[end] == -1 && start != end { // Must make sure a path is possible
		fmt.Println("No path exists...")
		return
	}
	if start == end || end == -1 {
		fmt.Printf("%v", start)
	} else {
		g.traversePath(start, g.Parent[end])
		fmt.Printf(" %v", end)
	}
}

// BreadthFirstSearch performs a vanilla BFS traversal of the graph using the
// Traversal struct as is GraphProcessor
func (g *Graph) BreadthFirstSearch(start int) {
	t := new(Traversal)
	g.bfs(start, t)
}

// Traversal implements the GraphProcessor interface in such a way to allow
// simple BFS and DFS traversals
type Traversal struct {
}

func (t *Traversal) processVertexEarly(g *Graph, v int) {
	fmt.Printf("Processed Vertex %v\n", v) // For Bfs & Dfs
}

func (t *Traversal) processVertexLate(g *Graph, v int) {
	// Do nothing here
}

func (t *Traversal) processEdge(g *Graph, x int, y int) {
	fmt.Printf("Processed Edge (%v, %v)\n", x, y) // For Bfs & Dfs
}

// ConnectedComponentTraversal implements the GraphProcessor interface so as to
// identify connected components while employing BFS
type ConnectedComponentTraversal struct {
}

func (t *ConnectedComponentTraversal) processVertexEarly(g *Graph, v int) {
	fmt.Printf(" %d", v)
}

func (t *ConnectedComponentTraversal) processVertexLate(g *Graph, v int) {
	// Do nothing here
}

func (t *ConnectedComponentTraversal) processEdge(g *Graph, x int, y int) {
	// Do nothing here
}

// ConnectedComponents discovers all connected components of a graph
func (g *Graph) ConnectedComponents() {
	t := new(ConnectedComponentTraversal)
	g.InitSearch()
	c := 1 // component number
	for i := 1; i <= g.nVertices; i++ {
		if g.State[i] == UNDISCOVERED {
			fmt.Printf("Component %v:", c)
			g.bfs(i, t)
			fmt.Println("")
			c++
		}
	}
}

// CycleFindTraversal implements GraphProcessor in order to find graph cycles
// with the help of DFS
type CycleFindTraversal struct {
}

func (t *CycleFindTraversal) processVertexEarly(g *Graph, v int) {
	// Do nothing here
}

func (t *CycleFindTraversal) processVertexLate(g *Graph, v int) {
	// Do nothing here
}

func (t *CycleFindTraversal) processEdge(g *Graph, x int, y int) {
	if g.Parent[y] != x { // Found back edge
		fmt.Printf("Cycle exists from %v to %v \n", y, x)
		fmt.Printf("Path is: ")
		g.FindPath(y, x)
		fmt.Printf("\n\n")
		g.Finished = true
	}
}

// FindCycles figures out if there are any cycles in the graph (nodes which connect in a cyclic fashion)
func (g *Graph) FindCycles(start int) {
	t := new(CycleFindTraversal)
	g.dfs(start, t)
}

// ArticulationVectorTraversal implements the interface GraphProcessor in order to
// find articulator vectors using DFS
type ArticulationVectorTraversal struct {
}

func (t *ArticulationVectorTraversal) processVertexEarly(g *Graph, v int) {
	g.ReachableAncestor[v] = v
}

func (t *ArticulationVectorTraversal) processEdge(g *Graph, x int, y int) {
	class := g.edgeClassification(x, y)
	if class == TREE {
		g.TreeOutDegree[x]++
	}
	if class == BACK && g.Parent[x] != y {
		if g.EntryTime[y] < g.EntryTime[g.ReachableAncestor[x]] {
			g.ReachableAncestor[x] = y
		}
	}
}

func (t *ArticulationVectorTraversal) processVertexLate(g *Graph, v int) {
	if g.Parent[v] == -1 { // Test if v is root
		if g.TreeOutDegree[v] > 1 { // root has more then one child
			fmt.Println("Root articulation vertex: ", v)
		}
		return
	}
	root := (g.Parent[g.Parent[v]] < 1) // Is the parent of v the root vertex?
	if g.ReachableAncestor[v] == g.Parent[v] && !root {
		// fmt.Printf("\n Reachable Ancestor: %v of %v with parent: %v \n", g.ReachableAncestor[v], v, g.Parent[v])
		fmt.Println("Parent Articulation Vector: ", g.Parent[v])
	}
	if g.ReachableAncestor[v] == v {
		// fmt.Println("Bridge Articulation Vertex: ", g.Parent[v])
		if g.TreeOutDegree[v] > 0 { // Check that v is not a leaf
			fmt.Println("Bridge Articulation Vertex: ", v)
		}
	}

	timeV := g.EntryTime[g.ReachableAncestor[v]]
	timeParent := g.EntryTime[g.ReachableAncestor[g.Parent[v]]]
	if timeV < timeParent {
		g.ReachableAncestor[g.Parent[v]] = g.ReachableAncestor[v]
	}
}

// FindArticulationVectors finds all the articulator vectors in a Graph
func (g *Graph) FindArticulationVectors(start int) {
	t := new(ArticulationVectorTraversal)
	g.dfs(start, t)
}

// DepthFirstSearch performs a DFS from a starting graph vertice
func (g *Graph) DepthFirstSearch(start int) {
	t := new(Traversal)
	g.dfs(start, t)
}

// dfs performs a general purpose Depth-first search through a graph
// processing each iteration as per the passed in GraphProcessor
// Implicitly uses a stack (via recursion) to prioritize node discovery
func (g *Graph) dfs(start int, p GraphProcessor) {

	if g.Finished {
		return
	}

	g.State[start] = DISCOVERED
	g.Time++
	g.EntryTime[start] = g.Time
	p.processVertexEarly(g, start)

	edgeNode := g.Edges[start]
	for edgeNode != nil {
		y := edgeNode.Y
		if g.State[y] == UNDISCOVERED {
			g.Parent[y] = start
			p.processEdge(g, start, y)
			g.dfs(y, p)
		} else if g.State[y] != PROCESSED || g.Directed {
			p.processEdge(g, start, y)
		}
		if g.Finished {
			return
		}
		edgeNode = edgeNode.Next
	}
	p.processVertexLate(g, start)
	g.Time++
	g.ExitTime[start] = g.Time
	g.State[start] = PROCESSED

}

// edgeClassification classifies an edge into a TREE, BACK, FORWARD or CROSS edge
func (g *Graph) edgeClassification(x, y int) EdgeType {
	if g.Parent[y] == x {
		return TREE
	}
	if g.State[y] == DISCOVERED {
		return BACK
	} else if g.State[y] == PROCESSED && (g.EntryTime[y] > g.EntryTime[x]) {
		return FORWARD
	} else if g.State[y] == PROCESSED && (g.EntryTime[y] < g.EntryTime[x]) {
		return CROSS
	}
	panic("Unclassified Edge")
}

var edgeTypes = [...]string{
	"TREE",
	"BACK",
	"FORWARD",
	"CROSS",
}

// String for the EdgeType enables this enum to appear as a string when passed to fmt
func (edgeType EdgeType) String() string {
	return edgeTypes[edgeType-1]
}

var states = [...]string{
	"UNDISCOVERED",
	"DISCOVERED",
	"PROCESSED",
}

// String for the VerticeState enables this enum to appear as a string when passed to fmt
func (state VerticeState) String() string {
	return states[state-1]
}
