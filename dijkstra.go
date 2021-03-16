package go_dijkstra

import "math"

type shortestPathFinder struct {
	source  int
	distTo  []int
	edgesTo []edge
}

func (s *shortestPathFinder) distanceTo(w int) int {
	if w < 0 || w > len(s.distTo) {
		return -1
	}

	return s.distTo[w]
}

func (s *shortestPathFinder) pathTo(w int) []edge {
	path := make([]edge, 0)
	var firstEdge int
	for e := s.edgesTo[w]; e.from() != s.source; e = s.edgesTo[e.from()] {
		firstEdge = e.from()
		path = append(path, e)
	}
	path = append(path, s.edgesTo[firstEdge])
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func newShortestPathFinder(source int, g *graph) shortestPathFinder {
	queue := &indexedMinQueue{
		data:        make([]keyVal, 0),
		dataToIndex: make(map[keyVal]int),
		keyMap:      make(map[int]struct{}),
	}

	visited := make([]bool, g.vertexCount())
	distTo := make([]int, g.vertexCount())
	edges := make([]edge, g.vertexCount())
	for i := 0; i < g.vertexCount(); i++ {
		distTo[i] = math.MaxInt64
	}
	distTo[source] = 0
	queue.insert(keyVal{
		key: source,
		val: 0,
	})
	for !queue.isEmpty() {
		v := queue.deleteMin()
		if visited[v.key] {
			continue
		}
		visited[v.key] = true
		for _, edge := range g.adjacent(v.key) {
			if edge.isEmpty() {
				continue
			}
			relax(queue, edge, distTo, edges)
		}
	}
	return shortestPathFinder{
		source:  source,
		distTo:  distTo,
		edgesTo: edges,
	}
}

func relax(queue *indexedMinQueue, e edge, distTo []int, edges []edge) {
	v, w := e.from(), e.to()
	if distTo[w] > distTo[v]+e.weight() {
		distTo[w] = distTo[v] + e.weight()
		edges[w] = e
	}
	kv := keyVal{
		key: w,
		val: distTo[w],
	}
	if queue.contains(kv) {
		queue.decreaseKey(kv, distTo[w])
	} else {
		queue.insert(kv)
	}
}
