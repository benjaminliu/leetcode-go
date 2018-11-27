package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	totalLen := m + n

	idx1, idx2, idx := m-1, n-1, totalLen-1

	for idx1 >=0 && idx2 >= 0 {
		if nums1[idx1] > nums2[idx2] {
			nums1[idx] = nums1[idx1]
			idx1--
		} else {
			nums1[idx] = nums2[idx2]
			idx2--
		}

		idx--
	}
	if idx1 <0{
		for idx2 >=0{
			nums1[idx] =nums2[idx2]
			idx2--
			idx--
		}
	}


}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{4, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
