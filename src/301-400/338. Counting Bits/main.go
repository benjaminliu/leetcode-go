package main

import "fmt"

func countBits(num int) []int {
	if num < 0 {
		return nil
	}
	if num == 0 {
		return []int{0}
	}

	if num == 1 {
		return []int{0, 1}
	}
	dp := make([]int, num+1)
	dp[0] = 0
	dp[1] = 1
	base := 2
	var j int
	for i := 2; i <= num; i++ {
		if i == base {
			dp [i ] = 1
			base *= 2
			j = 1
		} else {
			dp[i] = 1 + dp[j]
			j++
		}
	}
	return dp
}

func main() {
	fmt.Println(countBits(16))
}
