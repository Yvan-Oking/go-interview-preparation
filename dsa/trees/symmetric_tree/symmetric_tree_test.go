package main

import "testing"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric returns true if the binary tree is symmetric around its center
func isSymmetric(root *TreeNode) bool {
	// TODO: Implement your solution here
	return false
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

func TestIsSymmetric(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Symmetric tree",
			input:    []int{1, 2, 2, 3, 4, 4, 3},
			expected: true,
		},
		{
			name:     "Non-symmetric tree",
			input:    []int{1, 2, 2, -1, 3, -1, 3},
			expected: false,
		},
		{
			name:     "Empty tree",
			input:    []int{},
			expected: true,
		},
		{
			name:     "Single node tree",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "Complex non-symmetric tree",
			input:    []int{1, 2, 2, 2, -1, 2, -1},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := createTree(tc.input)
			result := isSymmetric(root)
			if result != tc.expected {
				t.Errorf("isSymmetric(%v) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
