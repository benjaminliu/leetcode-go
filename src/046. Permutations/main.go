package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	if length == 1 {
		return append(res, nums)
	}
	group := make([]int, length)
	visited := make([]byte, length)
	res = helper(nums, visited, 0, length-1, group, res)
	return res
}
func helper(nums []int, visited []byte, idx int, last int, group []int, res [][]int) [][]int {

	for i := 0; i <= last; i++ {
		if visited[i] == 1 {
			continue
		}
		group[idx] = nums[i]
		if idx == last {
			cp := make([]int, last+1)
			copy(cp, group)
			res = append(res, cp)
			return res
		}
		visited[i] = 1
		res = helper(nums, visited, idx+1, last, group, res)
		visited[i] = 0
	}
	return res
}

func permute1(nums []int) [][]int {
	res := make([][]int, 0)
	helper1(nums, 0, len(nums), &res)
	return res
}
func helper1(nums []int, idx int, length int, res *[][]int) {
	if idx == len(nums) {
		cp := make([]int, length)
		copy(cp, nums)
		*res = append(*res, cp)
		return
	}

	for i := idx; i < length; i++ {
		swap(nums, idx, i)
		helper1(nums, idx+1, length, res)
		swap(nums, idx, i)
	}
}
func swap(nums []int, i int, j int) {
	if i != j {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute1(nums))
}
