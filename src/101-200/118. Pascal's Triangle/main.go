package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}
	res := make([][]int, 0)
	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}
	for i := 1; i < numRows; i++ {
		group := make([]int, i+1)
		last := res[i-1]
		group[0] = 1
		for j := 1; j < i; j++ {
			group[j] = last[j-1] + last[j]
		}
		group[i] = 1
		res = append(res, group)
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
