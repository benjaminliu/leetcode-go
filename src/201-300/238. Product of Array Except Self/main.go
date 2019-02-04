package main

import "fmt"

func productExceptSelf(nums []int) []int {
	len1 := len(nums)
	if len1 == 0 {
		return nil
	}

	if len1 == 1 {
		return nums
	}

	target := make([]int, len1)

	for i := 0; i < len1; i++ {
		target[i] = 1
	}

	for i := 0; i < len1; i++ {
		for j := 0; j < len1; j++ {
			if i == j {
				continue
			}
			target[i] *= nums[j]
		}
	}
	return target
}

func productExceptSelf1(nums []int) []int {
	len1 := len(nums)
	if len1 < 2 {
		return nums
	}

	left := make([]int, len1)
	right := make([]int, len1)
	for i := 0; i < len1; i++ {
		left[i] = 1
		right[i] = 1
	}
	left[0] = nums[0]
	for i := 1; i < len1; i++ {
		left[i] = left[i-1] * nums[i]
	}
	last := len1 - 1
	right[last] = nums[last]
	for i := last - 1; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}

	nums[0] = right[1]
	nums[last] = left[last-1]
	for i := 1; i < last; i++ {
		nums[i] = left[i-1] * right[i+1]
	}
	return nums
}

func productExceptSelf2(nums []int) []int {
	len1 := len(nums)
	if len1 < 2 {
		return nums
	}

	res := make([]int, len1)
	res[0] = 1
	for i := 1; i < len1; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	right := 1
	for i := len1 - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf1(nums))
}
