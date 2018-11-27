package main

import "sort"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	if length == 1 {
		return append(res, nums)
	}
	sort.Ints(nums)
	group := make([]int, length)
	visited := make([]byte, length)
	res = helper(nums, visited, 0, length-1, group, res)
	return res
}
func helper(nums []int, visited []byte, idx int, last int, group []int, res [][]int) [][]int {
	if idx == last +1{
		return res
	}

	for i := 0; i <= last; i++ {
		if visited[i] == 1 {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && visited[i-1] == 0 {
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

func main() {

}
