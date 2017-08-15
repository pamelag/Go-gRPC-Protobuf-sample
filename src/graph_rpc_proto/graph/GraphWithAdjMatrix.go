package graph

import "fmt"

const MAX_VERTICES int = 20

type Queue struct {
	values []int
}

func (q Queue) Insert(index int) {
	q.values = append(q.values, index)
}

func (q Queue) Remove() int {
	size := len(q.values)
	removed := q.values[0]
	q.values = q.values[1:size]
	return removed
}

func NewQueue() Queue {
	return Queue{values: make([]int, 0)}
}


type Stack struct {
	values []int
}
func (s Stack) push(item int) {
	s.values = append(s.values, item)
}
func (s Stack) pop()  {
	size := len(s.values)
	s.values = s.values[size-1:size]
}

func (s Stack) size() int {
	return len(s.values)
}

func NewStack() Stack {
	return Stack{values: make([]int, 0)}
}

type Vertex struct {
	Label string
	WasVisited bool
	Neighbours []Vertex
}

func NewVertex(label int, wasVisited bool) Vertex {
	neighbours := make([]Vertex, 0)
	vertex := Vertex{Label: "0", WasVisited: false, Neighbours: neighbours}
	return vertex
}

type Graph struct {
	Name string
	MaxVertices int
	VertexList []Vertex
	AdjacencyMatrix map[int][]int
	NoOfVertex int
	PrevNoOfVertex int
	SStack Stack
	Queue Queue
}

func (g *Graph) SetName(name string) {
	g.Name = name
}

func (g *Graph) GetName() string {
	return g.Name
}

func (g *Graph) AddVertex(label int, visited bool) {
	g.PrevNoOfVertex = len(g.VertexList)
	vertex := NewVertex(label, visited)
	g.VertexList = append(g.VertexList, vertex)
	g.NoOfVertex = len(g.VertexList)
	// update adj matrix
	g.AdjacencyMatrix[label] = make([]int, g.NoOfVertex)

}

func (g *Graph) AddEdge(start int, end int, lStart string, lEnd string) {
	startV := Vertex{}
	endV := Vertex{}
	g.AdjacencyMatrix[start][end] = 1
	g.AdjacencyMatrix[end][start] = 1
	for i:=0; i < len(g.VertexList); i++ {
		if g.VertexList[i].Label == lStart {
			startV = g.VertexList[i]
		}
		if g.VertexList[i].Label == lEnd {
			endV = g.VertexList[i]
		}
	}
	startV.Neighbours = append(startV.Neighbours, endV)
}

func (g *Graph) updateAdjMatrixOnGraphResize() {
	//noOfVertices := len(g.VertexList)
	//newAdjacencyMatrix := [noOfVertices][noOfVertices]int{}
	//newAdjacencyMatrix := make(map[int]int, noOfVertices)
	//for i := 0; i <  g.NoOfVertex; i++ {
	//	for p := 0; p < noOfVertices; p++  {
	//		if i > g.PrevNoOfVertex {
	//			//newAdjacencyMatrix[i][p] = 0
	//			lst := newAdjacencyMatrix[i]
	//			lst = append(lst, 0)
	//		}
	//		//newAdjacencyMatrix[i][p] = g.AdjacencyMatrix[i][p]
	//		lst := newAdjacencyMatrix[i]
	//		lst = append(lst, p)
	//	}
	//}
}

func NewGraphStruct() Graph {
	vertexList := make([]Vertex, 20)
	vertex := Vertex{Label: "0", WasVisited: false}

	vertexList[0]=vertex
	AdjacencyMatrix := make(map[int][]int, 0)
	graph := Graph{VertexList:vertexList, AdjacencyMatrix:AdjacencyMatrix, SStack: NewStack(), Queue : NewQueue()}

	return graph
}

func (g *Graph) PrintVertices() {
	for i := 0; i < g.NoOfVertex; i++ {
		fmt.Println(g.VertexList[i])
	}
}

func (g *Graph) PrintVertex(i int) {
	fmt.Print(g.VertexList[i]);
}

func (g *Graph) DFS() {
	//root := g.VertexList[0]
	g.VertexList[0].WasVisited = true
	g.PrintVertex(0)
	g.SStack.push(1)

	for g.SStack.size() < 1 {
		unVisited :=  g.GetAdjUnvisitedVertex(0)

		if unVisited == -1 {
			g.SStack.pop()
		} else {
			g.VertexList[unVisited].WasVisited = true
			//g.PrevNoOfVertex(0)
			g.SStack.push(unVisited)
		}

		for j := 0; j < g.NoOfVertex; j++  {
			g.VertexList[j].WasVisited = false
		}
	}
}

func (g *Graph) bfs() {
	root := g.VertexList[0]
	root.WasVisited = true
	g.PrintVertex(0)
	g.Queue.Insert(0)
	v2 := 0

	for g.SStack.size() < 1 {
		v1 := g.Queue.Remove()
		// until it has no unvisited neighbours
		v2=g.GetAdjUnvisitedVertex(v1)
		for v2 != -1 {
			g.VertexList[v2].WasVisited = true
			g.PrintVertex(v2)
			g.Queue.Insert(v2)
		}
	}

	for j:=0; j<g.NoOfVertex; j++ {
		g.VertexList[j].WasVisited = false
	}

}

func (g *Graph) GetAdjUnvisitedVertex(v int) int {
	unVisited := 0
	for i := 0; i < g.NoOfVertex; i++ {
		if g.AdjacencyMatrix[v][i] == 1 && g.VertexList[i].WasVisited == false {
			unVisited = i
		} else {
			unVisited = -1
		}
	}
	return unVisited
}
