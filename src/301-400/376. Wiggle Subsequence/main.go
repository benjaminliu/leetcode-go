package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	if n == 2 {
		if nums[0] == nums[1] {
			return 1
		}
		return 2
	}
	diff := make([]int, n)

	idx := 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			diff[idx] = 1
			idx++
		} else if nums[i] < nums[i-1] {
			diff[idx] = -1
			idx++
		} else {
		}
	}
	if idx == 0 {
		return 1
	}
	res := 2
	for i := 1; i < idx; i++ {
		if diff[i] == diff[i-1] {
			continue
		}
		res++
	}

	return res
}

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(wiggleMaxLength([]int{0, 0, 0}))
}
