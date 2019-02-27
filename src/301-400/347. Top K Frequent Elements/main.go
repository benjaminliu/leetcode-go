package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	if n == 1 {
		return nums
	}

	maps := make(map[int]int)

	for _, value := range nums {
		maps[value] ++
	}

	map1 := make(map[int][]int)
	list := make([]int, 0)
	for key, value := range maps {
		if value1, ok := map1[value]; ok {
			value1 = append(value1, key)
			map1[value] = value1
		} else {
			map1[value] = []int{key}
			list = append(list, value)
		}
	}

	//bucket sort is good choice
	sort.Ints(list)

	res := make([]int, 0)
	count := 0
	for i := len(list) - 1; i >= 0; i-- {
		temp, _ := map1[list[i]]
		for _, t1 := range temp {
			res = append(res, t1)
			count++
			if count == k {
				return res
			}
		}
	}
	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
