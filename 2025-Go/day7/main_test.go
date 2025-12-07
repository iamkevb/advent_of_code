package main

import (
	"fmt"
	"testing"
)

func TestCountTimelines(t *testing.T) {
	tests := []struct {
		grid     []string
		expected int64
	}{
		{
			grid: []string{
				".......^.......",
				"...............",
			},
			expected: 2,
		},
		{
			grid: []string{
				".......^.......",
				"......^.^......",
			},
			expected: 4,
		},
		{
			grid: []string{
				".......^.......",
				"......^.^......",
				".....^.^.^.....",
			},
			expected: 8,
		},
		{
			grid: []string{
				".......^.......",
				"......^.^......",
				".....^.^.^.....",
				"....^.^...^....",
			},
			expected: 13,
		},
	}

	for i, tt := range tests {
		name := fmt.Sprintf("Test: %d", i)
		t.Run(name, func(t *testing.T) {
			grid := toRunes(tt.grid)
			timelines := make([]int64, len(grid[0]))
			start := len(tt.grid[0]) / 2
			timelines[start] = 1
			counts := countTimelines(grid, timelines)
			sum := int64(0)
			for _, v := range counts {
				sum += v
			}

			if tt.expected != sum {
				t.Errorf("%s: got %d; want %d", name, sum, tt.expected)
			}
		})
	}

}

func toRunes(strings []string) [][]rune {
	r := [][]rune{}
	for _, s := range strings {
		r = append(r, []rune(s))
	}
	return r
}
