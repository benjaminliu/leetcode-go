package main

import "fmt"

func removeDuplicates(nums []int) int {
	col := len(nums)
	if col <3 {
		return col
	}
	dup :=0
	for i :=2;i < col;i++{
		if nums[i] == nums[i-dup-2]{
			dup ++
			continue
		}
		nums[i-dup] = nums[i]
	}
	return col - dup
}

func main() {
	nums := []int{0,0,1,1,1,1,2,3,3}
	//nums := []int{1,1,1,2,2,3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
