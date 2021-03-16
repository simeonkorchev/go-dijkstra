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

	g = newGraph(8)
	g.addEdge(edge{
		v:   0,
		w:   1,
		wei: 5,
	})
	g.addEdge(edge{
		v:   0,
		w:   4,
		wei: 9,
	})
	g.addEdge(edge{
		v:   0,
		w:   7,
		wei: 8,
	})

	g.addEdge(edge{
		v:   4,
		w:   7,
		wei: 5,
	})
	g.addEdge(edge{
		v:   4,
		w:   5,
		wei: 4,
	})
	g.addEdge(edge{
		v:   4,
		w:   6,
		wei: 20,
	})
	g.addEdge(edge{
		v:   7,
		w:   2,
		wei: 7,
	})
	g.addEdge(edge{
		v:   7,
		w:   5,
		wei: 6,
	})
	g.addEdge(edge{
		v:   1,
		w:   7,
		wei: 4,
	})
	g.addEdge(edge{
		v:   1,
		w:   2,
		wei: 12,
	})
	g.addEdge(edge{
		v:   1,
		w:   3,
		wei: 15,
	})
	g.addEdge(edge{
		v:   2,
		w:   6,
		wei: 11,
	})
	g.addEdge(edge{
		v:   2,
		w:   3,
		wei: 3,
	})
	g.addEdge(edge{
		v:   5,
		w:   6,
		wei: 13,
	})
	g.addEdge(edge{
		v:   5,
		w:   2,
		wei: 1,
	})
	g.addEdge(edge{
		v:   3,
		w:   6,
		wei: 9,
	})
	pathFinder = newShortestPathFinder(0, g)
	dist := pathFinder.distanceTo(5)
	if dist != 13 {
		t.Errorf("expected distance to %d to be %d, but it is %d", 5, 5, dist)
	}
	dist = pathFinder.distanceTo(6)
	if dist != 25 {
		t.Errorf("expected distance to %d to be %d, but it is %d", 6, 25, dist)
	}
	if !reflect.DeepEqual(pathFinder.pathTo(6), []edge{
		{
			v:   0,
			w:   4,
			wei: 9,
		},
		{
			v:   4,
			w:   5,
			wei: 4,
		},
		{
			v:   5,
			w:   2,
			wei: 1,
		},
		{
			v:   2,
			w:   6,
			wei: 11,
		},
	}) {
		t.Fatalf("expected path did not match the actual")
	}
}
