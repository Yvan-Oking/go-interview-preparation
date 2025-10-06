package main

import (
	"reflect"
	"testing"
)

// mergeSort returns a sorted copy of the input slice using the merge sort algorithm
func mergeSort(nums []int) []int {
	// TODO: Implement your solution here
	return nil
}

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Normal case with unsorted array",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element array",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Already sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with duplicate elements",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of input to ensure it's not modified
			input := make([]int, len(tc.input))
			copy(input, tc.input)

			result := mergeSort(input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("mergeSort(%v) = %v; expected %v", tc.input, result, tc.expected)
			}

			// Check that original slice was not modified
			if !reflect.DeepEqual(input, tc.input) {
				t.Errorf("Input slice was modified: got %v; want %v", input, tc.input)
			}
		})
	}
}
