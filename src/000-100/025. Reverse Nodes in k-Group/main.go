package main

import "fmt"

func removeDuplicates(nums []int) int {
	length:=len(nums)
	if length < 2 {
		return length
	}

	dupCount :=0
	for i:=1;i < length;i++{
		if nums[i] == nums[i-1]{
			dupCount++
		}else if dupCount >0{
			nums[i-dupCount] = nums[i]
		}
	}
	return length-dupCount
}

func main() {
	nums :=[]int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(nums))
}
