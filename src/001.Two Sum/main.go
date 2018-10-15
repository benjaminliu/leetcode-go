package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int,2)
	m := make(map[int]int)
	for i :=0;i< len(nums);i++{
		temp := target - nums[i]
		if _,ok :=m[temp]; ok{
			res[0] = m[temp]
			res[1] = i
			break
		}
		m[nums[i]] = i
	}
	return res
}

func main() {
	res:=twoSum([]int{2,7,11,15} ,13)
	fmt.Println(res)
}
