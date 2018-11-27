package main

import "fmt"

func searchRange(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return []int{-1, -1}
	}

	if nums[0 ] > target {
		return []int{-1, -1}
	}

	last := length - 1

	if nums[last] < target {
		return []int{-1, -1}
	}

	start, end := -1, -1
	if nums[0] == target {
		start = 0
	} else {
		start = helperLeft(nums, target, 1, last)
	}

	if nums[last] == target {
		end = last
	} else {
		if start < 0 {
			end = helperRight(nums, target, 0, last-1)
		} else {
			end = helperRight(nums, target, start, last-1)
		}
	}

	return []int{start, end}
}

func helperRight(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2
	if nums[mid] > target {
		return helperRight(nums, target, left, mid-1)
	}

	if nums[mid] < target {
		return helperRight(nums, target, mid+1, right)
	}

	if nums[mid+1] != target {
		return mid
	}

	return helperRight(nums, target, mid+1, right)
}

func helperLeft(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2
	if nums[mid] > target {
		return helperLeft(nums, target, left, mid-1)
	}
	if nums[mid] < target {
		return helperLeft(nums, target, mid+1, right)
	}

	if nums[mid-1] != target {
		return mid
	}
	return helperLeft(nums, target, left, mid-1)
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}

	//fmt.Println(searchRange(nums,8))
	fmt.Println(searchRange(nums, 6))
}
