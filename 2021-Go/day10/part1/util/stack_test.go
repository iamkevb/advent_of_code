package util

import "testing"

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push('s')
	if stack.Empty() {
		t.Error("Stack should contain one element")
	}
}
