package main

import "testing"

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
	}

	for _, test := range tests {
		// Make a copy of the input array
		inputCopy := make([]int, len(test.input))
		copy(inputCopy, test.input)

		// Sort the input array using BubbleSort
		bubble_sort(inputCopy)

		// Compare the sorted array with the expected array
		for i := 0; i < len(inputCopy); i++ {
			if inputCopy[i] != test.expected[i] {
				t.Errorf("Test failed for input: %v. Expected: %v, Got: %v", test.input, test.expected, inputCopy)
				break
			}
		}
	}
}
