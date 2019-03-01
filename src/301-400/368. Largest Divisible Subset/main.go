package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return nums
	}

	sort.Ints(nums)

	dp := make([][]int, n)
	maxIdx, maxSize := 0, -1
	finalIdx, finalSize := 0, -1

	for i := 0; i < n; i++ {
		dp[i] = make([]int, 0)

		//nums[j] is the last num of dp[j], if cur nums is Divisible by the last num, it is divisible to and num (because next num is divisible to the previous num)
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if len(dp[j]) > maxSize {
					maxSize = len(dp[j])
					maxIdx = j
				}
			}
		}

		if maxSize != -1 {
			dp[i] = append(dp[i], dp[maxIdx]...)
		}
		dp[i] = append(dp[i], nums[i])

		if len(dp[i]) > finalSize {
			finalSize = len(dp[i])
			finalIdx = i
		}
		maxIdx = 0
		maxSize = -1
	}

	return dp[finalIdx]
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}))
}
