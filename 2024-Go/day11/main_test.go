package main

import "testing"

func Equal(arr1, arr2 []int64) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if arr2[i] != v {
			return false
		}
	}
	return true
}

func Test_Rule2(t *testing.T) {
	stones := rule2(11)
	if !Equal(stones, []int64{1, 1}) {
		t.Errorf("Expected 1,1 got %v", stones)
	}

	stones = rule2(123456)
	if !Equal(stones, []int64{123, 456}) {
		t.Errorf("Expected 1,1 got %v", stones)
	}

	stones = rule2(1006)
	if !Equal(stones, []int64{10, 6}) {
		t.Errorf("Expected 10,6 got %v", stones)
	}
}
