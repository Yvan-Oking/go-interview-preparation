# Coin Change

## Problem Description

You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.

You may assume that you have an infinite number of each kind of coin.

## Examples

**Example 1:**
```
Input: coins = [1,3,4], amount = 6
Output: 2
Explanation: 6 = 3 + 3 (using two 3-coin denominations)
```

**Example 2:**
```
Input: coins = [2], amount = 3
Output: -1
Explanation: The amount of 3 cannot be made up with coins of value 2.
```

**Example 3:**
```
Input: coins = [1], amount = 0
Output: 0
Explanation: No coins are needed to make up amount 0.
```

**Example 4:**
```
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1 (using one 1-coin and two 5-coin denominations)
```

## Constraints

- 1 <= coins.length <= 12
- 1 <= coins[i] <= 2^31 - 1
- 0 <= amount <= 10^4
