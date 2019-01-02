package main

import "fmt"

func rotate(nums []int, k int) {
	len1 := len(nums)
	if len1 < 2 {
		return
	}

	k = k % len1
	if k == 0 {
		return
	}

	tempArr := make([]int, k)
	for i := 0; i < k; i++ {
		tempArr[i] = nums[len1-k+i]
	}

	for i := 0; i < k; i ++ {
		last := nums[i]
		for j := k + i; j < len1; j += k {
			temp := nums[j ]
			nums[j] = last
			last = temp
		}
		nums[i] = tempArr[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
