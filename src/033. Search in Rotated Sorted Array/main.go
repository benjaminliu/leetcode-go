package main

import "fmt"

func search(nums []int, target int) int {
	length := len(nums)

	res := helper(nums, target, 0, length-1)
	return res
}
func helper(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		return mid
	}

	if nums[left] <= nums[mid] {
		if nums[left] <= target && nums[mid] > target {
			return helper(nums, target, left, mid-1)
		}
		return helper(nums, target, mid+1, right)
	}

	if nums[mid]  < target && nums[right] >= target {
		return helper(nums, target, mid+1, right)
	}

	return helper(nums, target, left, mid-1)
}

func main() {
	nums := []int{ 3,1}
	fmt.Println(search(nums,1))
}
