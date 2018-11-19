package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	len1 := len(nums)
	res := make([][]int, 0)
	res = append(res, make([]int, 0))
	if len1 == 0 {
		return res
	}

	sort.Ints(nums)

	group := make([]int, 0)
	for i := 0; i < len1; i++ {
		if (i > 0 && nums[i] == nums[i-1]) {
			continue
		}
		res = backtracking(nums, append(group, nums[i]), i+1, len1, res)
	}
	return res
}
func backtracking(nums []int, group []int, idx int, len1 int, res [][]int) [][]int {
	temp := make([]int, len(group))
	copy(temp,group)
	res = append(res, temp)
	if idx == len1 {
		return res
	}

	for i := idx; i < len1; i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}
		res = backtracking(nums, append(group, nums[i]), i+1, len1, res)
	}
	return res
}

func subsetsWithDup1(nums []int) [][]int {
	res := [][]int{}
	tmp := []int{}
	sort.Ints(nums)
	backTrack(&res, nums, tmp, 0)
	return res
}

func backTrack(res *[][]int, nums, tmp []int, start int) {
	*res = append(*res, append([]int(nil), tmp...))
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		t := append(tmp, nums[i])
		backTrack(res, nums, t, i+1)
	}
}

func main() {

}
