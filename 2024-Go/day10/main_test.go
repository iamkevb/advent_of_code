package main

import "testing"

func Test_CreateGridMap(t *testing.T) {
	input := [][]int{
		{9, 3, 9},
		{3, 2, 3},
		{5, 3, 9},
	}
	grid := createGridMap(input)
	if len(grid) != 9 {
		t.Errorf("Expected 9 elements in grid, got %d", len(grid))
	}
	idx := mapKey(1, 1)
	p, exists := grid[idx]
	if !exists {
		t.Errorf("Expected %s to exist", idx)
	}
	if p.col != 1 || p.row != 1 {
		t.Errorf("Expected x,y to be 1,1. Got: %d,%d", p.col, p.row)
	}
	if p.value != 2 {
		t.Errorf("Expected 2, got %d", p.value)
	}
	n0 := mapKey(1, 0)
	n1 := mapKey(1, 2)
	n2 := mapKey(0, 1)
	n3 := mapKey(2, 1)
	if len(p.next) != 4 || p.next[0] != n0 || p.next[1] != n1 || p.next[2] != n2 || p.next[3] != n3 {
		t.Errorf("next invalid : %v", p.next)
	}
}

func Test_FindStartingPositions(t *testing.T) {
	grid := map[string]Position{
		mapKey(0, 0): {0, 0, 0, []string{}},
		mapKey(1, 0): {0, 0, 5, []string{}},
		mapKey(1, 1): {0, 0, 4, []string{}},
		mapKey(0, 1): {0, 0, 0, []string{}},
	}
	starts := findStartingPositions(grid)
	if len(starts) != 2 {
		t.Errorf("Expected 2, got %d", len(starts))
	}
}

func TestFindTrails(t *testing.T) {
	input := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 1, 1},
		{1, 1, 2, 1, 1, 1, 1, 8, 1, 1},
		{2, 3, 4, 5, 6, 7, 8, 9, 1, 1},
	}
	grid := createGridMap(input)
	start := findStartingPositions(grid)
	score := findTrails(start[0], grid, 0)
	if score != 2 {
		t.Errorf("expected 2, but got %d", score)
	}
}
