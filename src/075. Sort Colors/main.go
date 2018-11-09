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

func main() {
	nums := []int{2,0,1}
	sortColors(nums)
	fmt.Println(nums)
}
