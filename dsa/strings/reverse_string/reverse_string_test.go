package main

import "testing"

// reverseString reverses the string in-place
func reverseString(s []byte) {
	// TODO: Implement your solution here
}

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Normal string",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "String with capital letters",
			input:    "Hannah",
			expected: "hannaH",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Two characters",
			input:    "ab",
			expected: "ba",
		},
		{
			name:     "Three characters",
			input:    "abc",
			expected: "cba",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert string to byte slice for testing
			s := []byte(tc.input)
			reverseString(s)
			result := string(s)
			if result != tc.expected {
				t.Errorf("reverseString(%s) = %s; expected %s", tc.input, result, tc.expected)
			}
		})
	}
}
