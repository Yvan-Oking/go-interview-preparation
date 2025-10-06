package main

import "testing"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor returns the lowest common ancestor of two nodes in a binary tree
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// TODO: Implement your solution here
	return nil
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

// findNode is a helper function to find a node with a given value in the tree
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	testCases := []struct {
		name     string
		tree     []int
		p        int
		q        int
		expected int
	}{
		{
			name:     "Normal case - nodes in different subtrees",
			tree:     []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			p:        5,
			q:        1,
			expected: 3,
		},
		{
			name:     "One node is ancestor of another",
			tree:     []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			p:        5,
			q:        4,
			expected: 5,
		},
		{
			name:     "Nodes are siblings",
			tree:     []int{1, 2, 3},
			p:        2,
			q:        3,
			expected: 1,
		},
		{
			name:     "Single path tree",
			tree:     []int{1, 2, -1, 3},
			p:        2,
			q:        3,
			expected: 2,
		},
		{
			name:     "Root is one of the nodes",
			tree:     []int{1, 2, 3},
			p:        1,
			q:        3,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := createTree(tc.tree)
			p := findNode(root, tc.p)
			q := findNode(root, tc.q)
			result := lowestCommonAncestor(root, p, q)
			if result.Val != tc.expected {
				t.Errorf("lowestCommonAncestor(tree=%v, p=%d, q=%d) = %d; expected %d",
					tc.tree, tc.p, tc.q, result.Val, tc.expected)
			}
		})
	}
}
