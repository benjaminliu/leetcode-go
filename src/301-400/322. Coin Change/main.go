package main

import (
	"fmt"
)

//dp
//from 1 cent to amount of cents,
//coin 1, 2, 5
//and amount is 20
//when we calc 20,  suppose we know 19, 18 and 15,  so 20 is min of three:  15 + 1 (add 5 cents), 18 + 1 (add 2 cents) , 19 + 1 (add 1 cent)
func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)

	for i := 1; i < max; i++ {
		dp[i] = max
	}

	n := len(coins)
	for i := 1; i <= amount; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
