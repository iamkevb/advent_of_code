package main

import "testing"

func TestSimpleFoldVertically(t *testing.T) {
	paper = Paper{}
	paper[Dot{x: 0, y: 0}] = true
	paper[Dot{x: 0, y: 2}] = true
	p := foldVertically(Fold{axis: "y", value: 1})
	if len(p) != 1 {
		t.Errorf("Expected 1 point, got %d", len(p))
	}
	if !p[Dot{x: 0, y: 0}] {
		t.Error("Dot at 0,0 should exist")
	}
}

func TestSimpleFoldHorizontally(t *testing.T) {
	paper = Paper{}
	paper[Dot{x: 0, y: 0}] = true
	paper[Dot{x: 2, y: 0}] = true
	p := foldHorizontally(Fold{axis: "y", value: 1})
	if len(p) != 1 {
		t.Errorf("Expected 1 point, got %d", len(p))
	}
	if !p[Dot{x: 0, y: 0}] {
		t.Error("Dot at 0,0 should exist")
	}
}
