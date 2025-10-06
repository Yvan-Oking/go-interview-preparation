package main

import "testing"

// canFinish returns true if it is possible to finish all courses
func canFinish(numCourses int, prerequisites [][]int) bool {
	// TODO: Implement your solution here
	return false
}

func TestCanFinish(t *testing.T) {
	testCases := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			name:          "No cycle",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      true,
		},
		{
			name:          "Has cycle",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expected:      false,
		},
		{
			name:          "No prerequisites",
			numCourses:    3,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			name:          "Complex graph without cycle",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      true,
		},
		{
			name:          "Complex graph with cycle",
			numCourses:    4,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
			expected:      false,
		},
		{
			name:          "Single course",
			numCourses:    1,
			prerequisites: [][]int{},
			expected:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canFinish(tc.numCourses, tc.prerequisites)
			if result != tc.expected {
				t.Errorf("canFinish(%d, %v) = %v; expected %v",
					tc.numCourses, tc.prerequisites, result, tc.expected)
			}
		})
	}
}
