package main

import (
	"fmt"
	"strconv"
)

func findMissingRanges(nums []int, lower, upper int) []string {
	group := make([]string, 0)
	len1 := len(nums)
	if len1 == 0 {
		return append(group, strconv.Itoa(lower)+"->"+strconv.Itoa(upper))
	}

	if nums[0] > lower {
		if nums[0]-lower == 1 {
			group = append(group, strconv.Itoa(lower))
		} else {
			group = append(group, strconv.Itoa(lower)+"->"+strconv.Itoa(nums[0]-1))
		}
	}

	for i := 1; i < len1; i++ {
		if nums[i]-nums[i-1] > 1 {
			helper(nums, i, &group)
		}
	}

	last := len1 - 1
	if nums[last ] < upper {
		if upper-nums[last] == 1 {
			group = append(group, strconv.Itoa(upper))
		} else {
			group = append(group, strconv.Itoa(nums[last]+1)+"->"+strconv.Itoa(upper))
		}
	}
	return group
}
func helper(nums []int, i int, group *[]string) {
	if nums[i]-nums[i-1] == 2 {
		*group = append(*group, strconv.Itoa(nums[i]-1))
	} else {
		*group = append(*group, strconv.Itoa(nums[i-1]+1)+"->"+strconv.Itoa(nums[i]-1))
	}
}

func main() {
	nums := []int{0, 1, 3, 50, 75}
	fmt.Println(findMissingRanges(nums, 0, 99))
}
