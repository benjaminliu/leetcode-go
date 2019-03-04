package main

import "math"

//dp, need more care
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	return helper(dp, 1, n)
}
func helper(dp [][]int, left int, right int) int {
	if left >= right {
		return 0
	}

	if dp[left][right] != 0 {
		return dp[left][right]
	}
	res := math.MaxInt32
	for i := left; i <= right; i++ {
		temp := i + max(helper(dp, left, i-1), helper(dp, i+1, right))
		res = min(res, temp)
	}
	dp[left][right] = res
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getMoneyAmount1(n int) int {
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for j := 2; j <= n; j++ {
		for i := j - 1; i > 0; i-- {
			globalMin := math.MaxInt32
			for k := i + 1; k < j; k++ {
				localMax := k + max(dp[i][k-1], dp[k+1][j])
				globalMin = min(globalMin, localMax)
			}
			if i+1 == j {
				dp[i][j] = i
			} else {
				dp[i][j] = globalMin
			}
		}
	}
	return dp[1][n]
}

func main() {

}
