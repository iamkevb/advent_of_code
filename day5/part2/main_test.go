package main

import (
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	text := "0,9 -> 5,9"
	reader := strings.NewReader(text)
	lines := readInput(reader)
	if len(lines) != 1 {
		t.Errorf("Expected one line, got %d", len(lines))
	}
	l := lines[0]
	if l.start.x != 0 || l.start.y != 9 || l.end.x != 5 || l.end.y != 9 {
		t.Error("Incorrect values")
	}
}
