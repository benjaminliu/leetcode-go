package main

import "fmt"

func numDecodings(s string) int {
	len1 := len(s)
	if len1 == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len1+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len1; i++ {
		cur := c2i(s[i-1])
		if cur > 0 {
			dp[i] = dp[i-1]
		}
		last := c2i(s[i-2])
		sum := last*10 + cur
		if sum < 27 && last != 0 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len1]
}

func c2i(c byte) int {
	return int(c - '0')
}

func main() {
	fmt.Println(numDecodings("12"))
}
