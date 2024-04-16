package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr      []int
		val      int
		expected int
	}{
		{[]int{1, 2, 3, 5, 6, 7, 8}, 4, -1}, // Value not found
		{[]int{1, 2, 3, 5, 6, 7, 8}, 1, 0},  // Value found at the beginning
		{[]int{1, 2, 3, 5, 6, 7, 8}, 8, 6},  // Value found at the end
	}

	for _, test := range tests {
		btSearched := BinarySearch(test.arr, test.val)
		if btSearched != test.expected {
			t.Errorf("For array %v and value %d, expected index %d, but got %d\n", test.arr, test.val, test.expected, btSearched)
		} else {
			fmt.Printf("Test passed. For array %v and value %d, expected index %d, got %d\n", test.arr, test.val, test.expected, btSearched)
		}
	}
}
