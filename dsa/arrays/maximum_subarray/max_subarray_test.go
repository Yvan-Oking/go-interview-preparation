package main

import "testing"

// maxSubArray returns the maximum sum of any contiguous subarray
func maxSubArray(nums []int) int {
	// TODO: Implement your solution here
	return 0
}

func TestMaxSubArray(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Mixed positive and negative numbers",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "Single positive number",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "All positive numbers",
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
		{
			name:     "Single negative number",
			nums:     []int{-1},
			expected: -1,
		},
		{
			name:     "Two negative numbers",
			nums:     []int{-2, -1},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxSubArray(tc.nums)
			if result != tc.expected {
				t.Errorf("maxSubArray(%v) = %d; expected %d", tc.nums, result, tc.expected)
			}
		})
	}
}
