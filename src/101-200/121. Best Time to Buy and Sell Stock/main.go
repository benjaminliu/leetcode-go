package main

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	len1 := len(prices)
	profit := 0
	low := prices[0]
	for i := 1; i < len1; i++ {
		cur := prices[i]
		if cur < low {
			low = cur
		} else {
			diff := cur - low
			if diff > profit {
				profit = diff
			}
		}
	}
	return profit
}

//dp
func maxProfit1(prices []int) int {
	len1 := len(prices)
	dp := make([]int, len1)

	maxSofar := 0
	for i := 0; i < len1; i++ {
		dp[i] = 0
		for j := 0; j < i; j++ {
			if prices[j] < prices[i] {
				dp[i] = maxInt(dp[i], prices[i]-prices[j])
			}
		}
		maxSofar = maxInt(dp[i], maxSofar)
	}
	return maxSofar
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {

}
