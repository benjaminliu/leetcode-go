package main

import "fmt"

func search(nums []int, target int) bool {
	col := len(nums)
	if col == 0 {
		return false
	}

	left, right := 0, col-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[left] {
			left = mid + 1
			continue
		}

		if nums[mid] > nums[right] {
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			right--
		}
	}
	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 3))
}
