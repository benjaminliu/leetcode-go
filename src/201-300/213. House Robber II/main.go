package main

import "fmt"

func rob(nums []int) int {
	len1 := len(nums)
	if len1 == 0 {
		return 0
	}
	if len1 == 1 {
		return nums[0]
	}
	if len1 == 2 {
		if nums[0] >= nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	num1 := realRob(nums, 0, len1-2)
	num2 := realRob(nums, 1, len1-1)

	if num1 >= num2 {
		return num1
	}
	return num2
}

func realRob(nums []int, start, end int) int {
	len1 := end - start + 1
	if len1 == 0 {
		return 0
	}
	if len1 == 1 {
		return nums[start]
	}
	if len1 == 2 {
		if nums[start] >= nums[start+1] {
			return nums[start]
		}
		return nums[start+1]
	}

	dp := make([]int, len1)

	dp[0] = nums[start]
	if nums[start] >= nums[start+1] {
		dp[1] = nums[start]
	} else {
		dp[1] = nums[start+1]
	}

	for i := 2; i < len1; i++ {
		temp := dp[i-2] + nums[start+i]
		if temp >= dp[i-1] {
			dp[i] = temp
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len1-1]
}

func main() {
	nums := []int{1, 3, 1, 3, 100}
	fmt.Println(rob(nums))
}
