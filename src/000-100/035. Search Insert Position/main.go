package main

import "fmt"

func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if nums[0] >= target {
		return 0
	}

	last := length - 1
	if nums[last] < target {
		return length
	}

	return helper(nums, target, 1, last)
}

func helper(nums []int, target int, left int, right int) int {
	if left >right {
		return left
	}

	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		return helper(nums, target, left, mid-1)
	}

	return helper(nums, target, mid+1, right)
}

func main() {
 	nums := []int { 1,2,4,6,7}

	fmt.Println(searchInsert(nums,3))
}
