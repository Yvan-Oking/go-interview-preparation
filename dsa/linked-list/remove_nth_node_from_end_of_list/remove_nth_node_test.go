package main

import (
	"reflect"
	"testing"
)

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd removes the nth node from the end of the list and returns the head
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// TODO: Implement your solution here
	return nil
}

// createLinkedList creates a linked list from a slice of integers
func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	return head
}

// linkedListToSlice converts a linked list to a slice for easier comparison
func linkedListToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "Remove last element",
			input:    []int{1, 2, 3, 4, 5},
			n:        1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Remove first element",
			input:    []int{1, 2, 3, 4, 5},
			n:        5,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "Remove middle element",
			input:    []int{1, 2, 3, 4, 5},
			n:        3,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "Single element list",
			input:    []int{1},
			n:        1,
			expected: []int{},
		},
		{
			name:     "Two element list, remove first",
			input:    []int{1, 2},
			n:        2,
			expected: []int{2},
		},
		{
			name:     "Two element list, remove last",
			input:    []int{1, 2},
			n:        1,
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := createLinkedList(tc.input)
			result := removeNthFromEnd(head, tc.n)
			resultSlice := linkedListToSlice(result)
			if !reflect.DeepEqual(resultSlice, tc.expected) {
				t.Errorf("removeNthFromEnd(%v, %d) = %v; expected %v",
					tc.input, tc.n, resultSlice, tc.expected)
			}
		})
	}
}
