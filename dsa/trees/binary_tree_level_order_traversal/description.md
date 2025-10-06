# Binary Tree Level Order Traversal

## Problem Description

Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

## Examples

**Example 1:**
```
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Explanation: 
Level 0: [3]
Level 1: [9, 20]
Level 2: [15, 7]
```

**Example 2:**
```
Input: root = [1]
Output: [[1]]
Explanation: There is only one node at level 0.
```

**Example 3:**
```
Input: root = []
Output: []
Explanation: The tree is empty.
```

## Constraints

- The number of nodes in the tree is in the range [0, 2000].
- -1000 <= Node.val <= 1000
