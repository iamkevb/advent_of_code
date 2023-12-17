package main

import (
	"testing"
)

func TestIsGears(t *testing.T) {
	input := []string{
		"3..",
		".*.",
		"..3",
	}
	n := isGear(1, input, 1)
	if n != 9 {
		t.Fatalf("Expect 9, got %d", n)
	}
}

func TestIsGears2(t *testing.T) {
	input := []string{
		"3..",
		".*.",
		"3..",
	}
	n := isGear(1, input, 1)
	if n != 9 {
		t.Fatalf("Expect 9, got %d", n)
	}
}

func TestIsGears3(t *testing.T) {
	input := []string{
		"..3",
		".*.",
		"..3",
	}
	n := isGear(1, input, 1)
	if n != 9 {
		t.Fatalf("Expect 9, got %d", n)
	}
}

func TestIsGears4(t *testing.T) {
	input := []string{
		"...",
		"3*3",
		"...",
	}
	n := isGear(1, input, 1)
	if n != 9 {
		t.Fatalf("Expect 9, got %d", n)
	}
}

func TestIsGears5(t *testing.T) {
	input := []string{
		"...",
		".*.",
		"3.3",
	}
	n := isGear(1, input, 1)
	if n != 9 {
		t.Fatalf("Expect 9, got %d", n)
	}
}

func TestIsGears6(t *testing.T) {
	input := []string{
		"3.3",
		".*.",
		"...",
	}
	n := isGear(1, input, 1)
	if n != 9 {
		t.Fatalf("Expect 9, got %d", n)
	}
}
