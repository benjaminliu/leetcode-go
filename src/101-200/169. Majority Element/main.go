package main

import "fmt"

func majorityElement(nums []int) int {
	slow, fast := 0, 1
	len1 := len(nums)
	deleted := make([]bool, len1)
	for fast < len1 {
		if nums[fast] != nums[slow] {
			deleted[fast] = true
			fast++
			deleted [slow] = true
			slow++
			for slow < fast && deleted[slow] == true {
				slow++
			}
			if slow == fast {
				fast++
			}
		} else {
			fast++
		}
	}

	return nums[slow]
}

func majorityElement1(nums []int) int {
	ret := nums[0]
	count := 0

	for _, num := range nums {
		if count == 0 {
			ret = num
		}

		if num == ret {
			count++
		} else {
			count--
		}
	}
	return ret
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
