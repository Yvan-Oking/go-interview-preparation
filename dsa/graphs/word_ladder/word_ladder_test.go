package main

import "testing"

// ladderLength returns the length of the shortest transformation sequence
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// TODO: Implement your solution here
	return 0
}

func TestLadderLength(t *testing.T) {
	testCases := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			name:      "Normal transformation sequence",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  5,
		},
		{
			name:      "End word not in list",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0,
		},
		{
			name:      "Single character difference",
			beginWord: "dog",
			endWord:   "fog",
			wordList:  []string{"fog"},
			expected:  2,
		},
		{
			name:      "No transformation possible",
			beginWord: "cat",
			endWord:   "dog",
			wordList:  []string{"bat", "rat", "hat"},
			expected:  0,
		},
		{
			name:      "Multiple possible paths",
			beginWord: "red",
			endWord:   "tax",
			wordList:  []string{"ted", "tex", "red", "tax", "tad", "rex"},
			expected:  4,
		},
		{
			name:      "Begin word equals end word",
			beginWord: "same",
			endWord:   "same",
			wordList:  []string{"same"},
			expected:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ladderLength(tc.beginWord, tc.endWord, tc.wordList)
			if result != tc.expected {
				t.Errorf("ladderLength(%q, %q, %v) = %d; expected %d",
					tc.beginWord, tc.endWord, tc.wordList, result, tc.expected)
			}
		})
	}
}
