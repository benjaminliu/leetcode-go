package main

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			//from 0 to i - 1 , each one can be root,
			//left child has dp[j], right child has dp[i-j-1], so one root is dp[j]*dp[i-j-1]
			dp[i] = dp[i] + dp[j]*dp[i-j-1]
		}
	}
	return dp[n]
}

func main() {

}
