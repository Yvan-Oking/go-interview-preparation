package main

import (
	"reflect"
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList reverses a singly linked list
func reverseList(head *ListNode) *ListNode {
	// TODO: Implement your solution here
	return nil
}

// Helper function to create a linked list from slice
func createList(vals []int) *ListNode {
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

// Helper function to convert linked list to slice
func listToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func TestReverseList(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Normal linked list",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Two element list",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element list",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := createList(tc.input)
			result := reverseList(head)
			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tc.expected) {
				t.Errorf("reverseList(%v) = %v; expected %v", tc.input, resultSlice, tc.expected)
			}
		})
	}
}
