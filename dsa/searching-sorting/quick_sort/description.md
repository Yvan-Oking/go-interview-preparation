# Quick Sort

## Problem Description

Implement the Quick Sort algorithm to sort an array of integers in ascending order.

Quick Sort is a divide-and-conquer algorithm that works by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays, according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively.

## Examples

**Example 1:**
```
Input: [64, 34, 25, 12, 22, 11, 90]
Output: [11, 12, 22, 25, 34, 64, 90]
Explanation: The array is sorted in ascending order using Quick Sort
```

**Example 2:**
```
Input: [5, 2, 8, 1, 9]
Output: [1, 2, 5, 8, 9]
Explanation: The array is sorted in ascending order
```

**Example 3:**
```
Input: [1]
Output: [1]
Explanation: Single element array is already sorted
```

## Time Complexity

- **Best Case**: O(n log n)
- **Average Case**: O(n log n)
- **Worst Case**: O(n²)

## Space Complexity

- **O(log n)** - Space required for recursion stack (average case)
- **O(n)** - Space required for recursion stack (worst case)

## Constraints

- 1 <= arr.length <= 10^4
- -10^4 <= arr[i] <= 10^4
