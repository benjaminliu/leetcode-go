package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	total := len1 + len2
	mid := total / 2
	var median, previous, i, j, idx int
	for {
		var num1Reached = i == len1
		var num2Reached = j == len2
		if num1Reached {
			median = nums2[j]
			j++
		} else if num2Reached {
			median = nums1[i]
			i++
		} else {
			if nums1[i] > nums2[j] {
				median = nums2[j]
				j++
			} else {
				median = nums1[i]
				i++
			}
		}

		if idx == mid {
			break
		}
		idx++
		previous = median
	}
	if (len1+len2)%2 == 0 {
		return float64(median+previous) / 2.0
	}
	return float64(median)
}

func main() {
	a := []int{1, 3}
	b := []int{2}
	c := findMedianSortedArrays(a, b)
	fmt.Println(c)
}
