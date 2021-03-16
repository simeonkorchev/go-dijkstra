package go_dijkstra

type edge struct {
	v   int
	w   int
	wei int
}

func (e *edge) from() int {
	return e.v
}

func (e *edge) to() int {
	return e.w
}

func (e *edge) weight() int {
	return e.wei
}

func (e *edge) isEmpty() bool {
	return e.wei == 0 && e.from() == 0 && e.to() == 0
}

type graph struct {
	adjList [][]edge
}

func newGraph(vertexCount int) *graph {
	g := graph{adjList: make([][]edge, vertexCount)}
	for i := 0; i < vertexCount; i++ {
		g.adjList[i] = make([]edge, vertexCount)
	}

	return &g
}

func (g *graph) addEdge(e edge) {
	g.adjList[e.from()] = append(g.adjList[e.from()], e)
}

func (g *graph) adjacent(v int) []edge {
	return g.adjList[v]
}

func (g *graph) vertexCount() int {
	return len(g.adjList)
}
