package main

import "fmt"

func removeElement(nums []int, val int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		for left <= right && nums[right] == val {
			right--
		}
		for left <= right && nums[left] != val {
			left++
		}

		if left > right {
			break
		}

		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return right + 1
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
}
