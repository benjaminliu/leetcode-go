package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	len1 := len(numbers)
	if len1 == 0 {
		return nil
	}
	last := len1 - 1
	for i := 0; i < last; i++ {
		temp := target - numbers[i]

		right := binarySearch(numbers, temp, i+1, last)
		if right != -1 {
			return []int{i + 1, right + 1}
		}
	}
	return nil
}
func binarySearch(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
