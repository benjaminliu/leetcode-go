package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	last := n - 1
	lastLast := last - 1

	for i := 0; i < lastLast; i++ {
		for j := i + 1; j < last; j++ {
			if nums[j] > nums[i] {
				for k := j + 1; k < n; k++ {
					if nums[k] > nums[j] {
						return true
					}
				}
			}
		}
	}
	return false
}

func increasingTriplet1(nums []int) bool {
	small, big := math.MaxInt32,math.MaxInt32

	for _, value := range nums {
		if value <= small { //smaller then small and big
			small = value
		} else if value <= big { //smaller then big
			big = value
		} else { //bigger then small and big
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet1([]int{5,4,3,2,1}))
}
