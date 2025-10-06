package main

import "testing"

// maxProfit returns the maximum profit from buying and selling stock
func maxProfit(prices []int) int {
	// TODO: Implement your solution here
	return 0
}

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Normal case with profit",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "No profit possible - decreasing prices",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Increasing prices - buy first, sell last",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Decreasing prices - no profit",
			prices:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Small array with profit",
			prices:   []int{2, 4, 1},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxProfit(tc.prices)
			if result != tc.expected {
				t.Errorf("maxProfit(%v) = %d; expected %d", tc.prices, result, tc.expected)
			}
		})
	}
}
