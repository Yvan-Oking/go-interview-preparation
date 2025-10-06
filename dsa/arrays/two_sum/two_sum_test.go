package main

import (
	"reflect"
	"testing"
)

// twoSum returns the indices of two numbers that add up to target
func twoSum(nums []int, target int) []int {
	// TODO: Implement your solution here
	return nil
}

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Basic case with two numbers",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Case with different indices",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Case with duplicate numbers",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("twoSum(%v, %d) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
			}
		})
	}
}
