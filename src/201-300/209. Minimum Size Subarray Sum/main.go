package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	len1 := len(nums)
	if len1 == 0 {
		return 0
	}

	right := 0
	sum := 0
	for right < len1 {
		sum += nums[right]
		if sum >= s {
			break
		}
		right++
	}
	if right == 0 {
		return 1
	}

	if right >= len1 {
		return 0
	}
	left := 0
	minLen := right - left + 1
	last := len1 - 1
outter:
	for left <= right && right < len1 {
		for sum >= s && left <= right {
			sum -= nums[left]
			left++
			if sum < s {
				break
			}
			calcMinLen(&minLen, left, right)
		}

		for sum < s && right < len1 {
			if right == last {
				break outter
			}
			right++
			sum += nums[right]
			if sum >= s {
				calcMinLen(&minLen, left, right)
			}
		}
	}
	return minLen
}

func calcMinLen(minLen *int, left int, right int) {
	temp := right - left + 1
	if temp < *minLen {
		*minLen = temp
	}
}

func minSubArrayLen1(s int, nums []int) int {
	left := 0
	sum := 0
	result := math.MaxInt64

	for index, num := range nums {
		sum += num

		for sum >= s {
			result = min(result, index+1-left)
			sum -= nums[left]
			left++

		}
	}

	if result != math.MaxInt64 {
		return result
	}

	return 0

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(minSubArrayLen(11, nums))
}
