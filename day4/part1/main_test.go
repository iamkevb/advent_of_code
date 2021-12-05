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
