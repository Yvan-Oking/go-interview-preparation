package main

import "testing"

// lengthOfLIS returns the length of the longest increasing subsequence
func lengthOfLIS(nums []int) int {
	// TODO: Implement your solution here
	return 0
}

func TestLengthOfLIS(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Basic increasing sequence",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{7, 6, 5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "Single element",
			nums:     []int{0},
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "All same elements",
			nums:     []int{7, 7, 7, 7, 7, 7, 7},
			expected: 1,
		},
		{
			name:     "Complex case",
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			name:     "Long sequence",
			nums:     []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := lengthOfLIS(tc.nums)
			if result != tc.expected {
				t.Errorf("lengthOfLIS(%v) = %d; expected %d", tc.nums, result, tc.expected)
			}
		})
	}
}
