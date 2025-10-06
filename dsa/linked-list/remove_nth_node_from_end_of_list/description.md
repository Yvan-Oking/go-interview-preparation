# Remove Nth Node From End of List

## Problem Description

Given the head of a linked list, remove the `nth` node from the end of the list and return its head.

## Examples

**Example 1:**
```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Explanation: We remove the 2nd node from the end (node with value 4).
```

**Example 2:**
```
Input: head = [1], n = 1
Output: []
Explanation: We remove the only node in the list.
```

**Example 3:**
```
Input: head = [1,2], n = 1
Output: [1]
Explanation: We remove the last node (node with value 2).
```

## Constraints

- The number of nodes in the list is `sz`.
- 1 <= sz <= 30
- 0 <= Node.val <= 100
- 1 <= n <= sz

## Follow-up

Could you do this in one pass?
