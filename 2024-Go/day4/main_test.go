package main

import (
	"testing"
)

func TestCountXMAS(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"AMXSXMAAMM",
		"XMASAMXAMM",
	}
	expected := []int{
		1,
		0,
		2,
	}
	for i, v := range input {
		result := countXMAS([]string{v})
		if expected[i] != result {
			t.Errorf("Expected %v, but got %d", expected, result)
		}
	}
}

func TestDiagonals1(t *testing.T) {
	input := []string{
		"ABC",
		"DEF",
		"GHI",
	}
	expected := []string{
		"A",
		"DB",
		"GEC",
		"HF",
		"I",
	}
	result := diagonalToRows(input)

	if len(result) != len(expected) {
		t.Errorf("Expected: %v, Received: %v", expected, result)
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected string %v, Received %v", expected[i], v)
		}
	}
}
