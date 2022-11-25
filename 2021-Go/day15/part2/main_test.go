package main

import (
	"testing"
)

func TestCreateGraph(t *testing.T) {
	tile := [][]int{{8}}
	expected := [][]int{
		{8, 9, 1, 2, 3},
		{9, 1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
	}
	graph := createGraph(tile)
	for x, r := range graph {
		for y, v := range r {
			if v != expected[x][y] {
				t.Errorf("expected [%d,%d] to be %d. Got %d", x, y, expected[x][y], v)
			}
		}
	}
}

func TestPriorityQueueInsert(t *testing.T) {
	p1 := Square{p: Point{x: 0, y: 0}, risk: 1}
	p2 := Square{p: Point{x: 1, y: 0}, risk: 2}
	p3 := Square{p: Point{x: 2, y: 0}, risk: 3}

	pq := PriorityQueue{}
	pq.insert(p3)
	pq.insert(p2)
	pq.insert(p1)

	if len(pq) != 3 {
		t.Errorf("len expect 3, got %d", len(pq))
	}
	if pq[0] != p1 || pq[1] != p2 || pq[2] != p3 {
		t.Error("failed")
	}
}

func TestPriorityQueueDequeue(t *testing.T) {
	p1 := Square{p: Point{x: 0, y: 0}, risk: 1}
	p2 := Square{p: Point{x: 1, y: 0}, risk: 2}
	p3 := Square{p: Point{x: 2, y: 0}, risk: 3}

	pq := PriorityQueue{}
	pq.insert(p3)
	pq.insert(p2)
	pq.insert(p1)

	p := pq.dequeue()
	if p != p1 {
		t.Error("wrong item dequeued")
	}
	if len(pq) != 2 {
		t.Errorf("expected len 2, but was %d", len(pq))
	}
}

func TestAdjacentSquares(t *testing.T) {
	sq := Square{p: Point{x: 2, y: 2}, risk: 8}
	graph := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	adj := adjacentSquares(sq, graph)
	if len(adj) != 2 {
		t.Errorf("expected len 2, got %d", len(adj))
	}
	if adj[0].p.x != 2 && adj[0].p.y != 1 && adj[0].risk != 5 {
		t.Errorf("expected {2,1},5, got {%d,%d}%d", adj[0].p.x, adj[0].p.y, adj[0].risk)
	}
	if adj[1].p.x != 1 && adj[1].p.y != 2 && adj[1].risk != 7 {
		t.Errorf("expected {1,2},7, got {%d,%d}%d", adj[1].p.x, adj[1].p.y, adj[1].risk)
	}
}
