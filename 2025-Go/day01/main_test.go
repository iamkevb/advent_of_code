package main

import "testing"

func TestPart2Random(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		start    int
		expected int
	}{
		{
			name:     "R5",
			input:    []int{5},
			start:    50,
			expected: 0,
		},
		{
			name:     "L5",
			input:    []int{-5},
			start:    50,
			expected: 0,
		},
		{
			name:     "L50",
			input:    []int{-50},
			start:    50,
			expected: 1,
		},
		{
			name:     "R50",
			input:    []int{50},
			start:    50,
			expected: 1,
		},
		{
			name:     "R1000",
			input:    []int{1000},
			start:    50,
			expected: 10,
		},
		{
			name:     "L68 from 50",
			input:    []int{-68},
			start:    50,
			expected: 1,
		},
		{
			name:     "R48 from 52",
			input:    []int{48},
			start:    52,
			expected: 1,
		},
		{
			name:     "L5 from 0",
			input:    []int{-5},
			start:    0,
			expected: 0,
		},
		{
			name:     "R5 from 0",
			input:    []int{5},
			start:    0,
			expected: 0,
		},
		{
			name:     "R100 from 0",
			input:    []int{100},
			start:    0,
			expected: 1,
		},
		{
			name:     "L100 from 0",
			input:    []int{-100},
			start:    0,
			expected: 1,
		},
		{
			name:     "L150 from 50",
			input:    []int{-150},
			start:    50,
			expected: 2,
		},
		{
			name:     "R150 from 50",
			input:    []int{150},
			start:    50,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := part2(tt.input, tt.start)

			if result != tt.expected {
				t.Errorf("part2(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPart2Around0(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		start    int
		expected int
	}{
		{
			name:     "Land on 0 going right",
			input:    []int{2},
			start:    98,
			expected: 1,
		},
		{
			name:     "Land on 0 going left",
			input:    []int{-2},
			start:    2,
			expected: 1,
		},
		{
			name:     "Pass 0 going right",
			input:    []int{2},
			start:    99,
			expected: 1,
		},
		{
			name:     "Pass 0 going left",
			input:    []int{-2},
			start:    1,
			expected: 1,
		},
		{
			name:     "Start on 0 going right",
			input:    []int{2},
			start:    0,
			expected: 0,
		},
		{
			name:     "Start on 0 going left",
			input:    []int{-2},
			start:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := part2(tt.input, tt.start)

			if result != tt.expected {
				t.Errorf("part2(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
