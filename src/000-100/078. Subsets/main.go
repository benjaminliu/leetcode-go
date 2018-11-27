package main

import "fmt"

func subsets(nums []int) [][]int {
	col := len(nums)
	res :=make([][]int,0)
	res=append(res, make([]int,0))
	if col ==0 {
		return res
	}
	group:= make([]int,0)
	for i:=0;i<col;i++{
		res= backtrack(nums,i+1,col, append(group, nums[i]),res)
	}
	return res
}
func backtrack(nums []int, idx int, col int, group []int, res [][]int) [][]int {
	temp := make([]int, len(group))
	copy(temp,group)
	res = append(res, temp)

	for i :=idx;i < col;i++{
		res = backtrack(nums,i+1,col, append(group, nums[i]),res)
	}
	return res
}

func main() {
	nums := []int{1,2,3}
	fmt.Println(subsets(nums))
}
