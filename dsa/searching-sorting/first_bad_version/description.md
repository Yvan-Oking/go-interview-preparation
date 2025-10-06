# First Bad Version

## Problem Description

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have `n` versions `[1, 2, ..., n]` and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API `bool isBadVersion(version)` which returns whether `version` is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

## Examples

**Example 1:**
```
Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.
```

**Example 2:**
```
Input: n = 1, bad = 1
Output: 1
Explanation: Version 1 is the first bad version.
```

**Example 3:**
```
Input: n = 2, bad = 2
Output: 2
Explanation: Version 2 is the first bad version.
```

## Constraints

- 1 <= bad <= n <= 2^31 - 1


## Time Complexity

- **O(log n)** - Binary search reduces the search space by half each iteration

## Space Complexity

- **O(1)** - Only using a constant amount of extra space

