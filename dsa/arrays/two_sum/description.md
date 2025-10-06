# Two Sum

## Problem Description

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

You can return the answer in any order.

## Examples

**Example 1:**
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**
```
Input: nums = [3,3], target = 6
Output: [0,1]
```

## Constraints

- 2 <= nums.length <= 10^4
- -10^9 <= nums[i] <= 10^9
- -10^9 <= target <= 10^9
- Only one valid answer exists.

## Hints

1. A really brute force way would be to search for all possible pairs of numbers but that would be too slow. Again, the thing that is tricky is that we have to return the indices, not the numbers.

2. So, what if we have a sorted array? Would that help us?

3. The second hint is to use a hash table. For each element, we can check if target - element exists in the hash table. If it does, we found our answer. If not, we add the current element to the hash table.
