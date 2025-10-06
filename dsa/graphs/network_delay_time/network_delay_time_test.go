package main

import "testing"

// networkDelayTime returns the time it takes for all nodes to receive the signal
func networkDelayTime(times [][]int, n int, k int) int {
	// TODO: Implement your solution here
	return 0
}

func TestNetworkDelayTime(t *testing.T) {
	testCases := []struct {
		name     string
		times    [][]int
		n        int
		k        int
		expected int
	}{
		{
			name: "Normal case with all nodes reachable",
			times: [][]int{
				{2, 1, 1},
				{2, 3, 1},
				{3, 4, 1},
			},
			n:        4,
			k:        2,
			expected: 2,
		},
		{
			name: "Some nodes unreachable",
			times: [][]int{
				{1, 2, 1},
				{2, 3, 2},
			},
			n:        4,
			k:        1,
			expected: -1,
		},
		{
			name:     "Single node",
			times:    [][]int{},
			n:        1,
			k:        1,
			expected: 0,
		},
		{
			name: "Multiple paths to same node",
			times: [][]int{
				{1, 2, 1},
				{2, 3, 2},
				{1, 3, 4},
			},
			n:        3,
			k:        1,
			expected: 3,
		},
		{
			name: "Cycle in graph",
			times: [][]int{
				{1, 2, 1},
				{2, 3, 2},
				{3, 1, 3},
			},
			n:        3,
			k:        1,
			expected: 3,
		},
		{
			name: "Different weights",
			times: [][]int{
				{1, 2, 9},
				{1, 3, 2},
				{3, 2, 3},
			},
			n:        3,
			k:        1,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := networkDelayTime(tc.times, tc.n, tc.k)
			if result != tc.expected {
				t.Errorf("networkDelayTime(%v, %d, %d) = %d; expected %d",
					tc.times, tc.n, tc.k, result, tc.expected)
			}
		})
	}
}
