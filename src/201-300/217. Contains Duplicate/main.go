package main

import "fmt"

func containsDuplicate(nums []int) bool {
	len1 := len(nums)
	if len1 < 2 {
		return false
	}

	maps := make(map[int]bool)

	for i := 0; i < len1; i++ {
		if _, ok := maps[nums[i]]; ok {
			return true
		}
		maps[nums[i]] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}
