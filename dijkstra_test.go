package go_dijkstra

import (
	"reflect"
	"testing"
)

func TestFindShortestPath(t *testing.T) {
	g := newGraph(5)
	g.addEdge(edge{
		v:   0,
		w:   1,
		wei: 5,
	})
	g.addEdge(edge{
		v:   0,
		w:   2,
		wei: 3,
	})
	g.addEdge(edge{
		v:   1,
		w:   3,
		wei: 1,
	})
	g.addEdge(edge{
		v:   2,
		w:   3,
		wei: 4,
	})
	g.addEdge(edge{
		v:   3,
		w:   4,
		wei: 4,
	})

	pathFinder := newShortestPathFinder(0, g)
	if pathFinder.distanceTo(4) != 10 {
		t.Fatalf("Expected distance to %d to be %d, but it is %d", 4, 10, pathFinder.distanceTo(4))
	}
	if !reflect.DeepEqual(pathFinder.pathTo(4), []edge{
		{
			v:   0,
			w:   1,
			wei: 5,
		},
		{
			v:   1,
			w:   3,
			wei: 1,
		},
		{
			v:   3,
			w:   4,
			wei: 4,
		},
	}) {
		t.Fatalf("expected path did not match the actual")
	}

}
