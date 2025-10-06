package main

import "testing"

// maxSubArray returns the sum of the contiguous subarray with the largest sum
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
			name:     "All positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			expected: -1,
		},
		{
			name:     "Single element array",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Array with one negative and one positive",
			nums:     []int{-2, 1},
			expected: 1,
		},
		{
			name:     "Large numbers",
			nums:     []int{1000, -1000, 1000, -1000, 1000},
			expected: 1000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of input to ensure it's not modified
			input := make([]int, len(tc.nums))
			copy(input, tc.nums)

			result := maxSubArray(input)
			if result != tc.expected {
				t.Errorf("maxSubArray(%v) = %d; expected %d", tc.nums, result, tc.expected)
			}

			// Check that original slice was not modified
			for i := range tc.nums {
				if input[i] != tc.nums[i] {
					t.Errorf("Input slice was modified at index %d: got %d; want %d",
						i, input[i], tc.nums[i])
				}
			}
		})
	}
}
