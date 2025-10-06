package main

import "testing"

// climbStairs returns the number of distinct ways to climb to the top
func climbStairs(n int) int {
	// TODO: Implement your solution here
	return 0
}

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Base case: n=1",
			input:    1,
			expected: 1,
		},
		{
			name:     "Base case: n=2",
			input:    2,
			expected: 2,
		},
		{
			name:     "Small value: n=3",
			input:    3,
			expected: 3,
		},
		{
			name:     "Medium value: n=4",
			input:    4,
			expected: 5,
		},
		{
			name:     "Medium value: n=5",
			input:    5,
			expected: 8,
		},
		{
			name:     "Larger value: n=10",
			input:    10,
			expected: 89,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := climbStairs(tc.input)
			if result != tc.expected {
				t.Errorf("climbStairs(%d) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}
