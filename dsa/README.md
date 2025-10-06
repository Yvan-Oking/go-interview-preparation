# Data Structures and Algorithms Study Guide

This is a comprehensive collection of Data Structures and Algorithms problems implemented in Go, designed to help you master fundamental computer science concepts and prepare for technical interviews.

## Content Structure

### Arrays
- [Two Sum](./arrays/two_sum/) - Find two numbers that add up to target
- [Best Time to Buy and Sell Stock](./arrays/best_time_to_buy_and_sell_stock/) - Maximum profit from stock trading
- [Container With Most Water](./arrays/container_with_most_water/) - Two-pointer technique for area maximization
- [Maximum Subarray](./arrays/maximum_subarray/) - Kadane's algorithm for subarray sum
- [Rotate Array](./arrays/rotate_array/) - Array rotation techniques

### Dynamic Programming
- [Fibonacci Number](./dynamic-programming/fibonacci_number/) - Classic DP problem with multiple approaches
- [Climbing Stairs](./dynamic-programming/climbing_stairs/) - Step counting with DP
- [Coin Change](./dynamic-programming/coin_change/) - Minimum coins for target amount
- [Longest Increasing Subsequence](./dynamic-programming/longest_increasing_subsequence/) - LIS algorithm
- [Maximum Subarray Sum](./dynamic-programming/maximum_subarray_sum/) - DP approach to subarray problems

### Graphs
- [Number of Islands](./graphs/number_of_islands/) - DFS/BFS for connected components
- [Clone Graph](./graphs/clone_graph/) - Graph traversal and deep copying
- [Course Schedule](./graphs/course_schedule/) - Topological sorting and cycle detection
- [Network Delay Time](./graphs/network_delay_time/) - Shortest path algorithms (Dijkstra)
- [Word Ladder](./graphs/word_ladder/) - BFS for shortest transformation path

### Linked List
- [Reverse Linked List](./linked-list/reverse_linked_list/) - Iterative and recursive approaches
- [Merge Two Sorted Lists](./linked-list/merge_two_sorted_lists/) - Merge sorted linked lists
- [Linked List Cycle](./linked-list/linked_list_cycle/) - Floyd's cycle detection algorithm
- [Remove Nth Node From End](./linked-list/remove_nth_node_from_end_of_list/) - Two-pointer technique
- [Intersection of Two Linked Lists](./linked-list/intersection_of_two_linked_lists/) - Finding common nodes

### Searching and Sorting
- [Binary Search](./searching-sorting/binary_search/) - Classic binary search implementation
- [First Bad Version](./searching-sorting/first_bad_version/) - Binary search variant
- [Search in Rotated Sorted Array](./searching-sorting/search_in_rotated_sorted_array/) - Modified binary search
- [Merge Sort](./searching-sorting/merge_sort/) - Divide and conquer sorting algorithm
- [Quick Sort](./searching-sorting/quick_sort/) - Efficient in-place sorting algorithm

### Strings
- [Reverse String](./strings/reverse_string/) - String manipulation techniques
- [Valid Palindrome](./strings/valid_palindrome/) - Two-pointer string validation
- [Longest Substring Without Repeating Characters](./strings/longest_substring_without_repeating_characters/) - Sliding window technique
- [Group Anagrams](./strings/group_anagrams/) - Hash map and sorting approaches
- [Longest Palindromic Substring](./strings/longest_palindromic_substring/) - Expand around centers technique

### Trees
- [Maximum Depth of Binary Tree](./trees/maximum_depth_of_binary_tree/) - Tree traversal and depth calculation
- [Symmetric Tree](./trees/symmetric_tree/) - Tree symmetry validation
- [Binary Tree Level Order Traversal](./trees/binary_tree_level_order_traversal/) - BFS for tree levels
- [Validate Binary Search Tree](./trees/validate_binary_search_tree/) - BST property validation
- [Lowest Common Ancestor](./trees/lowest_common_ancestor_of_binary_tree/) - LCA algorithms

## How to Use This Guide

Each problem folder contains:
- **`description.md`**: Detailed problem statement with examples and constraints
- **`[problem_name]_test.go`**: Comprehensive test cases following Go testing standards

### Running Tests

To run tests for a specific problem:

```bash
# Navigate to the problem directory
cd dsa/arrays/two_sum

# Run the tests
go test

# Run tests with verbose output
go test -v

# Run tests for a specific function
go test -run TestTwoSum
```

To run all tests in the repository:

```bash
# From the root directory
go test ./...

# Run all tests with verbose output
go test -v ./...
```

## Recommended Study Order

### Phase 1: Fundamentals
1. **Arrays**: Start with Two Sum, then move to more complex array problems
2. **Strings**: Master string manipulation and sliding window techniques
3. **Linked List**: Understand pointer manipulation and traversal

### Phase 2: Intermediate
4. **Trees**: Learn tree traversal (DFS, BFS) and tree properties
5. **Searching & Sorting**: Master binary search and sorting algorithms
6. **Dynamic Programming**: Start with Fibonacci, then progress to more complex DP

### Phase 3: Advanced
7. **Graphs**: Learn graph traversal and shortest path algorithms
8. **Complex DP**: Tackle more challenging dynamic programming problems

## Key Learning Objectives

### Algorithmic Techniques
- **Two Pointers**: Efficient array and string processing
- **Sliding Window**: Optimal substring and subarray problems
- **Hash Maps**: Fast lookups and frequency counting
- **Binary Search**: Efficient searching in sorted data
- **Dynamic Programming**: Breaking down complex problems
- **Graph Traversal**: DFS, BFS, and shortest path algorithms

### Data Structures
- **Arrays and Slices**: Fundamental collection types
- **Linked Lists**: Pointer-based data structures
- **Trees**: Hierarchical data organization
- **Graphs**: Network and relationship modeling
- **Hash Tables**: Fast key-value storage