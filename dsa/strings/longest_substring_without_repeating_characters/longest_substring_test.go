package main

import "testing"

// lengthOfLongestSubstring returns the length of the longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	// TODO: Implement your solution here
	return 0
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Normal case with repeating characters",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "All same characters",
			input:    "bbbbb",
			expected: 1,
		},
		{
			name:     "Multiple non-overlapping substrings",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: 1,
		},
		{
			name:     "Two different characters",
			input:    "au",
			expected: 2,
		},
		{
			name:     "Repeating pattern with longer substring",
			input:    "dvdf",
			expected: 3,
		},
		{
			name:     "Complex pattern",
			input:    "tmmzuxt",
			expected: 5,
		},
		{
			name:     "Special characters",
			input:    "!@#$%^&*()",
			expected: 10,
		},
		{
			name:     "Mixed alphanumeric and special",
			input:    "abc123!@#",
			expected: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.input)
			if result != tc.expected {
				t.Errorf("lengthOfLongestSubstring(%q) = %d; expected %d",
					tc.input, result, tc.expected)
			}
		})
	}
}
