package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	len1 := len(prices)
	var last int
	hasLast := false
	end := len1 - 1

	profit := 0
	for i := 0; i < end; i++ {
		if prices[i] < prices[i+1] {
			if hasLast == false {
				last = prices[i]
				hasLast = true
			}
		} else {
			if hasLast == true {
				profit += prices[i] - last
				hasLast = false
			}
		}
	}

	if hasLast {
		profit += prices[end] - last
	}
	return profit
}

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > 0 {
			profit += prices[i] - min
		}
		min = prices[i]
	}
	return profit
}

func main() {

}
