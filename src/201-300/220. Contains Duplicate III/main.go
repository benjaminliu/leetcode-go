package main

import (
	"fmt"
	"math"
)

// not understand
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k < 1 || t < 0 {
		return false
	}
	len1 := len(nums)

	maps := make(map[int]int)
	for i := 0; i < len1; i++ {
		remappedNum := nums[i] - math.MinInt32
		bucket := remappedNum / (t + 1)
		if _, ok := maps[bucket]; ok {
			return true
		}
		if value, ok := maps[bucket-1]; ok {
			if remappedNum-value <= t {
				return true
			}
		}
		if value, ok := maps[bucket+1]; ok {
			if value-remappedNum <= t {
				return true
			}
		}
		if len(maps) >= k {
			lastBucket := (nums[i-k] - math.MinInt32) / (t + 1)
			delete(maps, lastBucket)
		}
		maps[bucket] = remappedNum
	}
	return false
}

func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	len1 := len(nums)
	last := len1 - 1

	nt := t * -1

	for i := 0; i < last; i++ {
		end := i + k
		for j := i + 1; j <= end && j < len1; j++ {
			temp := nums[i] - nums[j]
			if (temp >= 0 && temp <= t) || (temp < 0 && temp >= nt) {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsNearbyAlmostDuplicate(nums, 3, 0))
	nums = []int{1, 0, 1, 1}
	fmt.Println(containsNearbyAlmostDuplicate(nums, 1, 2))
	nums = []int{1, 5, 9, 1, 5, 9}
	fmt.Println(containsNearbyAlmostDuplicate(nums, 2, 3))
}
