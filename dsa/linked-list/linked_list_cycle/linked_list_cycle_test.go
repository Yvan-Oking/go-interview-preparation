package main

import "testing"

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle returns true if the linked list has a cycle
func hasCycle(head *ListNode) bool {
	// TODO: Implement your solution here
	return false
}

// createLinkedList creates a linked list from a slice of integers
// If pos is >= 0, it creates a cycle by connecting the last node to the node at position pos
func createLinkedList(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	// Create nodes
	nodes := make([]*ListNode, len(vals))
	for i := range vals {
		nodes[i] = &ListNode{Val: vals[i]}
	}

	// Connect nodes
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// Create cycle if pos is valid
	if pos >= 0 && pos < len(nodes) {
		nodes[len(nodes)-1].Next = nodes[pos]
	}

	return nodes[0]
}

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		name     string
		values   []int
		pos      int
		expected bool
	}{
		{
			name:     "Cycle at beginning",
			values:   []int{3, 2, 0, -4},
			pos:      0,
			expected: true,
		},
		{
			name:     "Cycle in middle",
			values:   []int{1, 2},
			pos:      0,
			expected: true,
		},
		{
			name:     "No cycle",
			values:   []int{1},
			pos:      -1,
			expected: false,
		},
		{
			name:     "Empty list",
			values:   []int{},
			pos:      -1,
			expected: false,
		},
		{
			name:     "Long list with cycle",
			values:   []int{1, 2, 3, 4, 5, 6},
			pos:      2,
			expected: true,
		},
		{
			name:     "Long list without cycle",
			values:   []int{1, 2, 3, 4, 5, 6},
			pos:      -1,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := createLinkedList(tc.values, tc.pos)
			result := hasCycle(head)
			if result != tc.expected {
				t.Errorf("hasCycle(list with values=%v and cycle at pos=%d) = %v; expected %v",
					tc.values, tc.pos, result, tc.expected)
			}
		})
	}
}
