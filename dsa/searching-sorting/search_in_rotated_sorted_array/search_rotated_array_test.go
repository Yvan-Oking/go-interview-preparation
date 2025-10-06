package main

import "testing"

// search returns the index of target in nums, or -1 if it is not found
func search(nums []int, target int) int {
	// TODO: Implement your solution here
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
			name:     "Target exists in rotated array",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "Target not found in rotated array",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "Single element array - target found",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "Single element array - target not found",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		{
			name:     "No rotation - normal binary search",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "Target is first element after rotation",
			nums:     []int{3, 4, 5, 1, 2},
			target:   1,
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := search(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("search(%v, %d) = %d; expected %d",
					tc.nums, tc.target, result, tc.expected)
			}
		})
	}
}
