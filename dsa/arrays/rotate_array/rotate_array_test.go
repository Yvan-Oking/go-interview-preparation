package main

import (
	"reflect"
	"testing"
)

// rotate rotates the array to the right by k steps
func rotate(nums []int, k int) {
	// TODO: Implement your solution here
}

func TestRotate(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Normal rotation by 3 steps",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "Rotation with negative numbers",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			name:     "Small array rotation",
			nums:     []int{1, 2},
			k:        1,
			expected: []int{2, 1},
		},
		{
			name:     "No rotation needed",
			nums:     []int{1, 2, 3},
			k:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Single element array",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy for testing since we modify in place
			nums := make([]int, len(tc.nums))
			copy(nums, tc.nums)
			rotate(nums, tc.k)
			if !reflect.DeepEqual(nums, tc.expected) {
				t.Errorf("rotate(%v, %d) = %v; expected %v", tc.nums, tc.k, nums, tc.expected)
			}
		})
	}
}
