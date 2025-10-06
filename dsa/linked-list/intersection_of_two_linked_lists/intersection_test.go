package main

import "testing"

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode returns the node at which the two lists intersect
func getIntersectionNode(headA, headB *ListNode) *ListNode {
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

// createIntersectedLists creates two linked lists that intersect at a given position
func createIntersectedLists(listA, listB, intersection []int) (*ListNode, *ListNode) {
	// Create the intersection part first
	intersectionNode := createLinkedList(intersection)

	// If no intersection, return separate lists
	if len(intersection) == 0 {
		return createLinkedList(listA), createLinkedList(listB)
	}

	// Create listA
	headA := createLinkedList(listA)
	if headA == nil {
		return intersectionNode, createLinkedList(listB)
	}
	currentA := headA
	for currentA.Next != nil {
		currentA = currentA.Next
	}
	currentA.Next = intersectionNode

	// Create listB
	headB := createLinkedList(listB)
	if headB == nil {
		return headA, intersectionNode
	}
	currentB := headB
	for currentB.Next != nil {
		currentB = currentB.Next
	}
	currentB.Next = intersectionNode

	return headA, headB
}

func TestGetIntersectionNode(t *testing.T) {
	testCases := []struct {
		name         string
		listA        []int
		listB        []int
		intersection []int
		expected     int // The expected value at intersection point (-1 if no intersection)
	}{
		{
			name:         "Intersection in the middle",
			listA:        []int{4, 1},
			listB:        []int{5, 6, 1},
			intersection: []int{8, 4, 5},
			expected:     8,
		},
		{
			name:         "No intersection",
			listA:        []int{2, 6, 4},
			listB:        []int{1, 5},
			intersection: []int{},
			expected:     -1,
		},
		{
			name:         "Intersection at the beginning of one list",
			listA:        []int{},
			listB:        []int{3},
			intersection: []int{2, 4},
			expected:     2,
		},
		{
			name:         "Same length lists",
			listA:        []int{1},
			listB:        []int{2},
			intersection: []int{3, 4},
			expected:     3,
		},
		{
			name:         "Long lists with intersection",
			listA:        []int{1, 2, 3, 4},
			listB:        []int{5, 6},
			intersection: []int{7, 8, 9},
			expected:     7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			headA, headB := createIntersectedLists(tc.listA, tc.listB, tc.intersection)
			result := getIntersectionNode(headA, headB)

			if tc.expected == -1 {
				if result != nil {
					t.Errorf("getIntersectionNode() = node with value %d; expected nil", result.Val)
				}
			} else {
				if result == nil {
					t.Errorf("getIntersectionNode() = nil; expected node with value %d", tc.expected)
				} else if result.Val != tc.expected {
					t.Errorf("getIntersectionNode() = node with value %d; expected %d", result.Val, tc.expected)
				}
			}
		})
	}
}
