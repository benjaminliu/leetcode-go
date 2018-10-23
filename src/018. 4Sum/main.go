package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return [][]int{}
	}

	sort.Ints(nums)

	res := [][]int{}

	lastTwo := nums[length-2] + nums[length-1]

	lastThree := nums[length-3] + lastTwo

	for i, endi := 0, length-3; i < endi; i++ {
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+lastThree < target {
			continue
		}
		// remove duplication
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, endj := i+1, length-2; j < endj; j++ {
			twoSum := nums[i] + nums[j]
			if twoSum+nums[j+1]+nums[j+2] > target {
				break
			}
			if twoSum+lastTwo < target {
				continue
			}
			// remove duplication
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := length - 1
			for left < right {
				fSum := twoSum + nums[left] + nums[right]
				if fSum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					// remove duplication
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					// remove duplication
					for left < right && nums[right] == nums [right-1] {
						right--
					}

					left++
					right--
				} else if fSum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	fmt.Println(fourSum(nums, 0))
}
