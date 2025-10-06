package main

import "testing"

// coinChange returns the fewest number of coins needed to make up the amount
func coinChange(coins []int, amount int) int {
	// TODO: Implement your solution here
	return 0
}

func TestCoinChange(t *testing.T) {
	testCases := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Basic case with possible solution",
			coins:    []int{1, 3, 4},
			amount:   6,
			expected: 2,
		},
		{
			name:     "Case with no solution",
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			name:     "Amount is 0",
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			name:     "Single coin type",
			coins:    []int{1},
			amount:   1,
			expected: 1,
		},
		{
			name:     "Multiple coin types",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			name:     "Large amount",
			coins:    []int{1, 2, 5},
			amount:   100,
			expected: 20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := coinChange(tc.coins, tc.amount)
			if result != tc.expected {
				t.Errorf("coinChange(%v, %d) = %d; expected %d", tc.coins, tc.amount, result, tc.expected)
			}
		})
	}
}
