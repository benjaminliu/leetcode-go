package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	length := len(candidates)
	if length == 0 {
		return res
	}
	sort.Ints(candidates)
	group := make([]int, 0)
	res = helper(candidates, target, 0, length, group, res)
	return res
}
func helper(candidates []int, target int, idx int, length int, group []int, res [][]int) [][]int {
	for i := idx; i < length; i++ {
		if i > idx {
			if candidates[i] == candidates[i-1]{
				continue
			}
		}
		cur := candidates[i]
		if cur > target {
			return res
		}
		temp := target - cur
		if temp < 0 {
			return res
		}
		if temp == 0 {
			cp := make([]int, len(group))
			copy(cp, group)
			return append(res, append(cp, cur))
		}
		res = helper(candidates, temp, i+1, length, append(group, cur), res)
	}

	return res
}

func main() {
	var data = []int{2, 5, 2, 1, 2}
	fmt.Println(combinationSum2(data, 5))
}
