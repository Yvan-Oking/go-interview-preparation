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

// mergeTwoLists merges two sorted linked lists
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "Both lists have elements",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "First list is empty",
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
		{
			name:     "Second list is empty",
			list1:    []int{0},
			list2:    []int{},
			expected: []int{0},
		},
		{
			name:     "Both lists are empty",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element lists",
			list1:    []int{1},
			list2:    []int{2},
			expected: []int{1, 2},
		},
		{
			name:     "Different lengths",
			list1:    []int{1, 2, 3, 4, 5},
			list2:    []int{1, 2},
			expected: []int{1, 1, 2, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list1 := createList(tc.list1)
			list2 := createList(tc.list2)
			result := mergeTwoLists(list1, list2)
			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tc.expected) {
				t.Errorf("mergeTwoLists(%v, %v) = %v; expected %v", tc.list1, tc.list2, resultSlice, tc.expected)
			}
		})
	}
}
