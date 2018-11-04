package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] += i + 1
	}

	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = i * factorial[i-1]
	}

	res := ""
	k = k - 1
	for i := n; i > 0; i-- {
		idx := k / factorial[i-1]
		k = k % factorial[i-1]
		res += strconv.Itoa(nums[idx])
		nums = append(nums[:idx], nums[idx+1:] ...)
	}
	return res
}

func main() {
	fmt.Println(getPermutation(3, 3))
}
