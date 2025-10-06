package main

import (
	"reflect"
	"sort"
	"testing"
)

// groupAnagrams groups anagrams together
func groupAnagrams(strs []string) [][]string {
	// TODO: Implement your solution here
	return nil
}

// normalizeGroups sorts each group and the groups themselves for consistent comparison
func normalizeGroups(groups [][]string) [][]string {
	// Sort strings within each group
	for _, group := range groups {
		sort.Strings(group)
	}

	// Sort groups by their first string (or length if empty)
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) == 0 {
			return true
		}
		if len(groups[j]) == 0 {
			return false
		}
		return groups[i][0] < groups[j][0]
	})

	return groups
}

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "Multiple anagram groups",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:     "Single empty string",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Single character",
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "No anagrams",
			input:    []string{"a", "b"},
			expected: [][]string{{"a"}, {"b"}},
		},
		{
			name:     "One anagram pair",
			input:    []string{"ab", "ba"},
			expected: [][]string{{"ab", "ba"}},
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: [][]string{},
		},
		{
			name:     "Multiple characters same frequency",
			input:    []string{"aab", "aba", "baa", "abc"},
			expected: [][]string{{"aab", "aba", "baa"}, {"abc"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of input to ensure it's not modified
			input := make([]string, len(tc.input))
			copy(input, tc.input)

			result := groupAnagrams(input)

			// Normalize both result and expected for comparison
			normalizedResult := normalizeGroups(result)
			normalizedExpected := normalizeGroups(tc.expected)

			if !reflect.DeepEqual(normalizedResult, normalizedExpected) {
				t.Errorf("groupAnagrams(%v) = %v; expected %v",
					tc.input, result, tc.expected)
			}

			// Check that original slice was not modified
			if !reflect.DeepEqual(input, tc.input) {
				t.Errorf("Input slice was modified: got %v; want %v",
					input, tc.input)
			}
		})
	}
}
