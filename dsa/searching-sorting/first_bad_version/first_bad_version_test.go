package main

import "testing"

// isBadVersion is a helper function that simulates an API call
// It will be implemented differently for each test case
var isBadVersion func(version int) bool

// firstBadVersion returns the first bad version in a sequence of versions
func firstBadVersion(n int) int {
	// TODO: Implement your solution here
	return 0
}

func TestFirstBadVersion(t *testing.T) {
	testCases := []struct {
		name       string
		n          int
		badVersion int
		expected   int
	}{
		{
			name:       "First version is bad",
			n:          5,
			badVersion: 1,
			expected:   1,
		},
		{
			name:       "Last version is bad",
			n:          5,
			badVersion: 5,
			expected:   5,
		},
		{
			name:       "Bad version in middle",
			n:          5,
			badVersion: 3,
			expected:   3,
		},
		{
			name:       "Single version",
			n:          1,
			badVersion: 1,
			expected:   1,
		},
		{
			name:       "Large number of versions",
			n:          100,
			badVersion: 42,
			expected:   42,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the isBadVersion function for this test case
			isBadVersion = func(version int) bool {
				return version >= tc.badVersion
			}

			result := firstBadVersion(tc.n)
			if result != tc.expected {
				t.Errorf("firstBadVersion(%d) = %d; expected %d", tc.n, result, tc.expected)
			}
		})
	}
}
