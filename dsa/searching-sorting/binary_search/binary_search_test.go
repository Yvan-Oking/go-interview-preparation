package main

import (
	"math"
	"testing"
)

// search performs binary search on a sorted array
func search(nums []int, target int) int {
	// TODO: Implement your solution here
	left := 0
	right := len(nums)

	for left < right {
		middle := int(math.Floor(float64((left + right) / 2)))

		if nums[middle] == target {
			return middle
		}

		if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return -1
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Target found in middle",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "Target not found",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
		{
			name:     "Single element array - target found",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element array - target not found",
			nums:     []int{5},
			target:   -5,
			expected: -1,
		},
		{
			name:     "Two element array - target at end",
			nums:     []int{2, 5},
			target:   5,
			expected: 1,
		},
		{
			name:     "Two element array - target at start",
			nums:     []int{2, 5},
			target:   2,
			expected: 0,
		},
		{
			name:     "Two element array - target not found",
			nums:     []int{2, 5},
			target:   3,
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := search(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("search(%v, %d) = %d; expected %d", tc.nums, tc.target, result, tc.expected)
			}
		})
	}
}
