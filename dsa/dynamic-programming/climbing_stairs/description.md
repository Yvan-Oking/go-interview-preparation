# Climbing Stairs

## Problem Description

You are climbing a staircase. It takes `n` steps to reach the top.

Each time you can either climb `1` or `2` steps. In how many distinct ways can you climb to the top?

## Examples

**Example 1:**
```
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```

**Example 2:**
```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

**Example 3:**
```
Input: n = 4
Output: 5
Explanation: There are five ways to climb to the top.
1. 1 + 1 + 1 + 1
2. 1 + 1 + 2
3. 1 + 2 + 1
4. 2 + 1 + 1
5. 2 + 2
```

## Constraints

- 1 <= n <= 45


## Time Complexity

- **O(n)** - We compute each value from 1 to n exactly once

## Space Complexity

- **O(1)** - We only need to store the last two values (space-optimized approach)
- **O(n)** - If using a DP array to store all values
