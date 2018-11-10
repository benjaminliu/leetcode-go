package main

import "fmt"

func sortColors(nums []int) {
	col := len(nums)
	if col < 2 {
		return
	}
	i, j, k := -1, -1, -1

	for idx := 0; idx < col; idx++ {
		if nums[idx] == 0 {
			k++
			nums[k] = 2
			j++
			nums[j] = 1
			i++
			nums[i] = 0
		} else if nums[idx] == 1 {
			k++
			nums[k] = 2
			j++
			nums[j] = 1
		} else {
			k++
			nums[k] = 2
		}
	}
}

func sortColors1(nums []int) {
	col := len(nums)
	if col < 2 {
		return
	}

	left := 0
	for left < col {
		if nums[left] != 0 {
			break
		}
		left++
	}
	right := col - 1
	for left <= right {
		if nums[right] != 2 {
			break
		}
		right--
	}
	idx := left
	for idx <= right {
		if idx <left{
			idx = left
			continue
		}
		if nums[idx] == 0 {
			nums[idx], nums[left] = nums[left], nums[idx]
			left++
		} else if nums[idx] == 2 {
			nums[idx], nums[right] = nums[right], nums[idx]
			right--
		} else {
			idx ++
		}
	}
}

func main() {
	//nums := []int{2, 0, 1}
	nums := []int{2,   1}
	//nums := []int{2, 0, 2, 1, 1, 0}
	sortColors1(nums)

	fmt.Println(nums)
}
