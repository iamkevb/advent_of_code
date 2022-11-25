package main

import "testing"

func TestTop(t *testing.T) {
	cave := [][]int{
		{9, 1, 9},
		{9, 9, 9},
	}
	lowPoints := lowPoints(cave)
	if len(lowPoints) != 1 || lowPoints[0].r != 1 || lowPoints[0].c != 0 {
		t.Errorf("incorrect point returned %x\n", lowPoints)
	}
}
