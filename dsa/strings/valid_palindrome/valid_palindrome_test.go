package main

import "testing"

// isPalindrome checks if the string is a palindrome, considering only alphanumeric characters and ignoring case
func isPalindrome(s string) bool {
	// TODO: Implement your solution here
	return false
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Standard palindrome with spaces and punctuation",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Not a palindrome",
			input:    "race a car",
			expected: false,
		},
		{
			name:     "Single space",
			input:    " ",
			expected: true,
		},
		{
			name:     "Simple palindrome",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "Simple non-palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "Mixed case palindrome",
			input:    "Madam",
			expected: true,
		},
		{
			name:     "Complex palindrome with apostrophes",
			input:    "No 'x' in Nixon",
			expected: true,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "Special characters only",
			input:    ".,?!",
			expected: true,
		},
		{
			name:     "Numbers palindrome",
			input:    "12321",
			expected: true,
		},
		{
			name:     "Mixed alphanumeric palindrome",
			input:    "A1b2c2b1a",
			expected: true,
		},
		{
			name:     "Mixed alphanumeric non-palindrome",
			input:    "A1b2c3b1a",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("isPalindrome(%q) = %v; expected %v",
					tc.input, result, tc.expected)
			}
		})
	}
}
