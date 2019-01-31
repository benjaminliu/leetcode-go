package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	len1 := len(nums)
	if len1 == 0 {
		return nil
	}
	res := make([]string, 0)
	group := make([]byte, 0)
	if len1 == 1 {
		appendI(&group, nums[0])
		res = append(res, string(group))
		return res
	}
	start := 0
	appendI(&group, nums[0])

	for i := 1; i < len1; i++ {
		if nums[i]-nums[i-1] == 1 {
			continue
		}

		end := i - 1
		if end != start {
			group = append(group, '-')
			group = append(group, '>')
			appendI(&group, nums[i-1])
		}
		res = append(res, string(group))
		group = group[0:0]
		start = i
		appendI(&group, nums[i])
	}
	if nums[len1-1]-nums[len1-2] == 1 {
		group = append(group, '-')
		group = append(group, '>')
		appendI(&group, nums[len1-1])
	}
	res = append(res, string(group))
	return res
}

func appendI(group *[]byte, i int) {

	strconv.Itoa(i)
	*group = append(*group, strconv.Itoa(i)...)
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
	nums = []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
	nums = []int{-1}
	fmt.Println(summaryRanges(nums))
	nums = []int{-2147483648,-2147483647,2147483647}
	fmt.Println(summaryRanges(nums))
}
