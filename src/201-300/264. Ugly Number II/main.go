package main

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1

	factor2, factor3, factor5 := 2, 3, 5

	idx2, idx3, idx5 := 0, 0, 0

	for i := 1; i < n; i++ {
		min1 := min(min(factor2, factor3), factor5)
		dp[i] = min1

		//cannot use else
		//because when 6 = 2 * 3, we need to move both 2 and 3 idx
		if factor2 == min1 {
			idx2++
			factor2 = 2 * dp[idx2]
		}
		if factor3 == min1 {
			idx3++
			factor3 = 3 * dp[idx3]
		}
		if factor5 == min1 {
			idx5++
			factor5 = 5 * dp[idx5]
		}
	}
	return dp[n-1]
}

func min(left, right int) int {
	if left <= right {
		return left
	}
	return right
}

func main() {

}
