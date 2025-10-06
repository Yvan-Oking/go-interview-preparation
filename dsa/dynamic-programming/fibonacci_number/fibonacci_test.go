package main

import "testing"

// fib returns the nth Fibonacci number
func fib(n int) int {
	// TODO: Implement your solution here
	return 0
}

func TestFib(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Base case: n=0",
			input:    0,
			expected: 0,
		},
		{
			name:     "Base case: n=1",
			input:    1,
			expected: 1,
		},
		{
			name:     "Base case: n=2",
			input:    2,
			expected: 1,
		},
		{
			name:     "Small value: n=3",
			input:    3,
			expected: 2,
		},
		{
			name:     "Small value: n=4",
			input:    4,
			expected: 3,
		},
		{
			name:     "Small value: n=5",
			input:    5,
			expected: 5,
		},
		{
			name:     "Medium value: n=6",
			input:    6,
			expected: 8,
		},
		{
			name:     "Medium value: n=10",
			input:    10,
			expected: 55,
		},
		{
			name:     "Large value: n=30",
			input:    30,
			expected: 832040,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := fib(tc.input)
			if result != tc.expected {
				t.Errorf("fib(%d) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}
