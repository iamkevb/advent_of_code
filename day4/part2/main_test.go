package main

import "testing"

func TestToInts(t *testing.T) {
	input := "22 13 17 11  0"

	got := toInts(input, " ")

	if len(got) != 5 {
		t.Errorf("Expected array len 5, got %d", len(got))
		t.Error(got)
	}
}

func TestRowSolved(t *testing.T) {
	board := Board{
		squares: [][]int{
			{-1, -1, -1, -1, 99},
			{-1, -1, 15, -1, 19},
			{18, 8, -1, 26, 20},
			{22, -1, 13, 6, -1},
			{-1, -1, 12, 3, -1},
		},
	}

	board.mark(99)

	if !board.rowSolved() {
		t.Error("This should be solved")
	}

	if !board.rowSolved() {
		t.Error("This should be solved")
	}
	if !board.checkSolved() {
		t.Error("this should be solved")
	}
	if !board.checkSolved() {
		t.Error("this should be solved")
	}
}
