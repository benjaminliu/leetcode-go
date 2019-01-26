package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	t1 := 9 * k
	if n < k || n > t1 {
		return nil
	}

	res := make([][]int, 0)
	group := make([]int, 0)

	helper(&res, group, 1, k, n)

	return res
}
func helper(res *[][]int, group []int, start int, k int, n int) {
	if k == 0 {
		if n == 0 {
			addRes(group, res)
		}
		return
	}
	if n <= 0 {
		return
	}
	t2 := start * k
	if n < t2 {
		return
	}
	//if n == t2 {
	//	for i := k; i > 0; i-- {
	//		group = append(group, start)
	//	}
	//	addRes(group, res)
	//	return
	//}

	t1 := 9 * k
	if n > t1 {
		return
	}
	//if n == t1 {
	//	for i := k; i > 0; i-- {
	//		group = append(group, 9)
	//	}
	//	addRes(group, res)
	//	return
	//}

	for i := start; i < 10; i++ {
		helper(res, append(group, i), i+1, k-1, n-i)
	}
}

func addRes(group []int, res *[][]int) {
	tempGroup := make([]int, len(group))
	copy(tempGroup, group)
	*res = append((*res), tempGroup)
}

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
}
