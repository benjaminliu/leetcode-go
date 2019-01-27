package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	len1 := len(nums)
	maps := make(map[int]int)
	for i := 0; i < len1; i++ {
		if idx, ok := maps[nums[i]]; ok {
			if i-idx <= k {
				return true
			}
		}
		maps[nums[i]] = i
	}
	return false
}

func containsNearbyDuplicate1(nums []int, k int) bool {
	len1 := len(nums)
	last := len1 - 1

	for i := 0; i < last; i++ {
		end := i + k
		for j := i + 1; j <= end && j < len1; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsNearbyDuplicate(nums, 3))
	nums = []int{1, 0, 1, 1}
	fmt.Println(containsNearbyDuplicate(nums, 1))
	nums = []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}
