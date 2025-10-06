package main

import "testing"

// longestPalindrome returns the longest palindromic substring
func longestPalindrome(s string) string {
	// TODO: Implement your solution here
	return ""
}

// isPalindrome checks if a string is a palindrome
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		possible []string // Multiple possible answers
	}{
		{
			name:     "Multiple palindromes same length",
			input:    "babad",
			possible: []string{"bab", "aba"},
		},
		{
			name:     "Single palindrome",
			input:    "cbbd",
			possible: []string{"bb"},
		},
		{
			name:     "Single character",
			input:    "a",
			possible: []string{"a"},
		},
		{
			name:     "Two characters no palindrome",
			input:    "ac",
			possible: []string{"a", "c"},
		},
		{
			name:     "Odd length palindrome",
			input:    "racecar",
			possible: []string{"racecar"},
		},
		{
			name:     "No palindrome longer than 1",
			input:    "abcdef",
			possible: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name:     "Empty string",
			input:    "",
			possible: []string{""},
		},
		{
			name:     "All same characters",
			input:    "aaaa",
			possible: []string{"aaaa"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestPalindrome(tc.input)

			// Verify result is one of the possible answers
			valid := false
			for _, possible := range tc.possible {
				if result == possible {
					valid = true
					break
				}
			}

			if !valid {
				// If result is not in possible answers, check if it's a valid palindrome
				// and has the same length as expected
				if !isPalindrome(result) || len(result) != len(tc.possible[0]) {
					t.Errorf("longestPalindrome(%q) = %q; expected one of %v",
						tc.input, result, tc.possible)
				}
			}
		})
	}
}
