package main

import "testing"

func TestIsInvalidId2WithInvalidIds(t *testing.T) {
	tests := []string{
		"1212",
		"1111",
		"123123",
		"12121212",
	}
	for _, s := range tests {
		if !isInvalidId2(s) {
			t.Errorf("expected %s to be an invalid id", s)
		}
	}
}

func TestIsInvalidId2WithValidIds(t *testing.T) {
	tests := []string{
		"1234",
		"12312",
	}
	for _, s := range tests {
		if isInvalidId2(s) {
			t.Errorf("expected %s to be an valid id", s)
		}
	}
}
