package main

import "fmt"

func nextPermutation(nums []int) {
	length := len(nums)
	if length < 2 {
		return
	}

	idx := length - 1
	var target int
	for idx > 0 {
		if nums[idx] > nums[idx-1] {
			target = idx - 1
			break
		}
		idx--
	}
	//last permutation, go to first
	if idx == 0 {
		left, right := 0, length-1
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left ++
			right--
		}
		return
	}

	temp, end := nums[target], length-1
	//find the swap position that swap with target
	for i := end; i > target; i-- {
		if nums[i] > temp {
			nums[i], nums[target] = temp, nums[i]
			break
		}
	}
	target ++
	//revert after the target
	for target < end {
		nums[target], nums[end] = nums[end], nums[target]
		target++
		end --
	}
}

func main() {
	nums := []int{1, 5, 4, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}
