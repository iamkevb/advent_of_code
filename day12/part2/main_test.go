package main

import (
	"strings"
	"testing"
)

func TestCanVisit(t *testing.T) {
	path := []string{"start", "A", "A", "b", "b", "c"}
	if !canVisit(path, "d") {
		t.Error("Can visit any cave once")
	}
	if canVisit(path, "c") {
		t.Error("Can only visit a single small cave once.")
	}
	if canVisit(path, "b") {
		t.Error("Cannot visit lower case nodes more than twice")
	}
	if canVisit(path, "start") {
		t.Error("Cannot visit start or end more than once")
	}
	if !canVisit(path, "A") {
		t.Error("Can visit upper case nodes many times")
	}

	path = []string{"start", "A", "A", "b", "b", "c", "end"}
	if canVisit(path, "A") {
		t.Error("Cannot visit after end")
	}
}

func TestContains(t *testing.T) {
	arr := []string{"1", "2", "3"}
	if !contains(arr, "1") {
		t.Error("Array contains 1")
	}
	if contains(arr, "4") {
		t.Error("Array does not contain 4")
	}
}

func TestInvalidPath(t *testing.T) {
	path := []string{"start", "b", "d", "b", "A"}
	if canVisit(path, "b") {
		t.Error("Cannot visit any small cave 3 times")
	}
}
func TestValidPath(t *testing.T) {
	validPath := []string{"start", "A", "b", "A", "c", "A", "b", "A", "end"}
	path := []string{}
	for _, c := range validPath {
		if !canVisit(path, c) {
			t.Errorf("Failed to visit %s [%s]", c, strings.Join(path, ","))
		}
		path = append(path, c)
	}
}
