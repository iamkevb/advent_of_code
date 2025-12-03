package main

import "testing"

func TestMaxJoltage(t *testing.T) {
	tests := []struct {
		name      string
		bank      []int
		batteries int
		expected  int64
	}{
		{
			name:      "2 from 1234",
			bank:      []int{1, 2, 3, 4},
			batteries: 2,
			expected:  34,
		},
		{
			name:      "3 from 1234",
			bank:      []int{1, 2, 3, 4},
			batteries: 3,
			expected:  234,
		},
		{
			name:      "12 from 987654321111111",
			bank:      []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			batteries: 12,
			expected:  987654321111,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxJoltage(tt.batteries, tt.bank)

			if result != tt.expected {
				t.Errorf("%s: got %d; want %d", tt.name, result, tt.expected)
			}
		})
	}
}
