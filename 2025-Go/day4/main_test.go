package main

import "testing"

func TestMaxJoltage(t *testing.T) {
	tests := []struct {
		name     string
		row      int
		col      int
		expected int
	}{
		{
			name:     "0,0",
			row:      0,
			col:      0,
			expected: 0,
		},
		{
			name:     "0,2",
			row:      0,
			col:      2,
			expected: 1,
		},
		{
			name:     "1,0",
			row:      1,
			col:      0,
			expected: 1,
		},
		{
			name:     "2,6",
			row:      2,
			col:      6,
			expected: 1,
		},
		{
			name:     "4,9",
			row:      4,
			col:      9,
			expected: 1,
		},
		{
			name:     "9,0",
			row:      9,
			col:      0,
			expected: 1,
		},
		{
			name:     "9,8",
			row:      9,
			col:      8,
			expected: 1,
		},
	}

	grid := readInput("./input-test.txt")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isRoll(grid, tt.row, tt.col)

			if result != tt.expected {
				t.Errorf("%s: got %d; want %d", tt.name, result, tt.expected)
			}
		})
	}
}
