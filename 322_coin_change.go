package main

// func coinChange(coins []int, amount int) int {
// 	sort.Ints(coins)
// 	var rest = amount
// 	var num int
// 	if rest == 0 {
// 		return 0
// 	}

// 	for i := 0; i < len(coins); i++ {
// 		if rest < coins[len(coins)-i-1] {
// 			continue
// 		}
// 		for rest/coins[len(coins)-i-1] > 0 {
// 			rest -= coins[len(coins)-i-1]
// 			num++
// 		}
// 		if rest == 0 {
// 			return num
// 		}
// 	}

// 	return -1
// }

func coinChange(coins []int, amount int) int {
	// we initialize dp with amount + 1 because it's the worst case
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			// here we compare the current dp[j] with a new calculated value
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}

	// return -1 when no combination can sum to amount
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if a-c < 0 {
				continue
			}
			dp[a] = min(dp[a], dp[a-c]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

// You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

// Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

// You may assume that you have an infinite number of each kind of coin.

// Example 1:

// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
// Example 2:

// Input: coins = [2], amount = 3
// Output: -1
// Example 3:

// Input: coins = [1], amount = 0
// Output: 0

// Constraints:

// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 104
