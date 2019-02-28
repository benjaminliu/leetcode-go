package main

import "fmt"

var dp []int = []int{0, 10, 91}

//dp
// dp[0] = 1     -- only 0
// dp[1] = 10 + dp[0]    -- 0,1,2,3,4,5,6,7,8,9
// dp[2] = 9 * 9 + dp[1]  -- first digit 1-9,  second 0-9 except first digit
// dp[3] = 9 * 9 * 8 + dp[2]  -- first digit 1-9,  second 0-9 except first digit, third digit 0-9 except frist 2 digit
// dp[4] = 9 * 9 * 8 * 7 + dp[3]
//...
func countNumbersWithUniqueDigits(n int) int {
	if n < len(dp) {
		return dp[n]
	}

	for i := len(dp); i <= n; i++ {
		temp := 81
		for j := 2; j < i; j++ {
			temp *= 10 - j
		}
		dp = append(dp, temp+dp[i-1])
	}
	return dp[n]
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits(3))
}
