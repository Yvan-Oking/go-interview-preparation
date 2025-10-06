package main

import "testing"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBST validates if a binary tree is a valid BST
func isValidBST(root *TreeNode) bool {
	// TODO: Implement your solution here
	return false
}

// Helper function to create a binary tree from slice
func createTree(vals []int) *TreeNode {
	if len(vals) == 0 {
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

func TestIsValidBST(t *testing.T) {
	testCases := []struct {
		name     string
		vals     []int
		expected bool
	}{
		{
			name:     "Valid BST",
			vals:     []int{2, 1, 3},
			expected: true,
		},
		{
			name:     "Invalid BST - right child less than root",
			vals:     []int{5, 1, 4, -1, -1, 3, 6},
			expected: false,
		},
		{
			name:     "Single node",
			vals:     []int{1},
			expected: true,
		},
		{
			name:     "Empty tree",
			vals:     []int{},
			expected: true,
		},
		{
			name:     "Valid BST with negative values",
			vals:     []int{-10, -3, 0},
			expected: true,
		},
		{
			name:     "Invalid BST - left subtree has value greater than root",
			vals:     []int{10, 5, 15, -1, -1, 6, 20},
			expected: false,
		},
		{
			name:     "Complex valid BST",
			vals:     []int{8, 3, 10, 1, 6, -1, 14, -1, -1, 4, 7, 13},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := createTree(tc.vals)
			result := isValidBST(root)
			if result != tc.expected {
				t.Errorf("isValidBST(%v) = %v; expected %v", tc.vals, result, tc.expected)
			}
		})
	}
}
