package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	length := len(nums)

	sort.Ints(nums)

	minDiff := math.MaxInt32
	res := 0

	for idx, end := 0, length-2; idx < end; idx++ {
		temp := target - nums[idx]
		left, right := idx+1, length-1
		for left < right {
			diff := nums[left] + nums[right] - temp
			if diff == 0 {
				return target
			}
			if diff < 0 {
				if -diff < minDiff {
					minDiff = -diff
					res = target + diff
				}
				left++
			} else {
				if diff < minDiff {
					minDiff = diff
					res = target + diff
				}
				right--
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}
