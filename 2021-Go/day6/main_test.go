package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGrowFish(t *testing.T) {
	fish := []int{1, 1, 1, 1, 1, 1, 1, 1}
	// fish at [5] reproduce and go to stage[0]
	// fish at [7] are maturing babies and go to stage [0]
	// babies start at [6]
	expected := []int{2, 1, 1, 1, 1, 1, 1, 1}
	got := growFish(fish)
	assertEqualSlices(t, got, expected)
}

func assertEqualSlices(t *testing.T, got, expected []int) {
	passed := len(got) == len(expected)
	if passed {
		for i := 0; passed && i < len(got); i++ {
			passed = got[i] == expected[i]
		}
	}
	if !passed {
		t.Errorf("Expected: %s\nGot: %s", sliceToString(expected), sliceToString(got))
	}
}

func sliceToString(slice []int) string {
	s := []string{}
	for _, v := range slice {
		s = append(s, fmt.Sprintf("%d", v))
	}
	return strings.Join(s, ",")
}
