package main

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	len1 := len(prices)
	profit:=0
	low:=prices[0]
	for i := 1; i < len1; i++ {
		cur := prices[i]
		if cur < low {
			low = cur
		}else {
			diff := cur -low
			if diff > profit {
				profit = diff
			}
		}
	}
	return profit
}

func main() {

}
