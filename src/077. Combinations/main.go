package main

import "fmt"

func combine1(n int, k int) [][]int {
	if n == 0 || k == 0 || n < k {
		return [][]int{}
	}

	res := make([][]int, 0)
	group := make([]int, k)

	res = backtrack1(n, k, 1, 0, group, res)
	return res
}
func backtrack1(n int, k int, i int, position int, group []int, res [][]int) [][]int {
	if position == k {
		temp := make([]int, len(group))
		copy(temp, group)
		return append(res, temp)
	}
	for idx := i; idx <= n; idx++ {
		group[position] = idx
		res = backtrack1(n, k, idx+1, position+1, group, res)
	}
	return res
}

func combine2(n int, k int) [][]int {
	if n == 0 || k == 0 || n < k {
		return [][]int{}
	}

	res := make([][]int, 0)
	group := make([]int, 0)

	res = backtrack2(n, k, 1, group, res)
	return res
}
func backtrack2(n int, k int, i int, group []int, res [][]int) [][]int {
	if len(group) == k {
		temp := make([]int, len(group))
		copy(temp, group)
		return append(res, temp)
	}
	for idx := i; idx <= n; idx++ {
		res = backtrack2(n, k, idx+1, append(group, idx), res)
	}
	return res
}

func main() {
	fmt.Println(combine1(5, 4))
}
