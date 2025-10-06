package main

import "testing"

// numIslands returns the number of islands in the grid
func numIslands(grid [][]byte) int {
	// TODO: Implement your solution here
	return 0
}

func TestNumIslands(t *testing.T) {
	testCases := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "Single island",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "Multiple islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			name: "Diagonal islands",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			expected: 5,
		},
		{
			name: "No islands",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "Single cell island",
			grid: [][]byte{
				{'1'},
			},
			expected: 1,
		},
		{
			name:     "Empty grid",
			grid:     [][]byte{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the grid to ensure it's not modified
			gridCopy := make([][]byte, len(tc.grid))
			for i := range tc.grid {
				gridCopy[i] = make([]byte, len(tc.grid[i]))
				copy(gridCopy[i], tc.grid[i])
			}

			result := numIslands(gridCopy)
			if result != tc.expected {
				t.Errorf("numIslands(%v) = %d; expected %d", tc.grid, result, tc.expected)
			}

			// Check that original grid was not modified
			for i := range tc.grid {
				for j := range tc.grid[i] {
					if gridCopy[i][j] != tc.grid[i][j] {
						t.Errorf("Grid was modified at position [%d][%d]: got %c; want %c",
							i, j, gridCopy[i][j], tc.grid[i][j])
					}
				}
			}
		})
	}
}
