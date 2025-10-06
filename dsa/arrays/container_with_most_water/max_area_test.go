package main

import "testing"

// maxArea returns the maximum area of water that can be contained
func maxArea(height []int) int {
	// TODO: Implement your solution here
	return 0
}

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Complex case with multiple peaks",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Simple case with two equal heights",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "Case with decreasing then increasing heights",
			height:   []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			name:     "Small array with peak in middle",
			height:   []int{1, 2, 1},
			expected: 2,
		},
		{
			name:     "Ascending heights",
			height:   []int{1, 2, 4, 3},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxArea(tc.height)
			if result != tc.expected {
				t.Errorf("maxArea(%v) = %d; expected %d", tc.height, result, tc.expected)
			}
		})
	}
}
