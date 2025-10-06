package main

import "testing"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth returns the maximum depth of a binary tree
func maxDepth(root *TreeNode) int {
	// TODO: Implement your solution here
	return 0
}

// createTree is a helper function to create a binary tree from a slice
func createTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) && vals[i] != -1 {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(vals) && vals[i] != -1 {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Normal case with balanced tree",
			input:    []int{3, 9, 20, -1, -1, 15, 7},
			expected: 3,
		},
		{
			name:     "Right skewed tree",
			input:    []int{1, -1, 2},
			expected: 2,
		},
		{
			name:     "Empty tree",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Single node tree",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "Complete binary tree",
			input:    []int{1, 2, 3, 4, 5},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := createTree(tc.input)
			result := maxDepth(root)
			if result != tc.expected {
				t.Errorf("maxDepth(%v) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}
