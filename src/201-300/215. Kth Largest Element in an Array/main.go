package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	len1 := len(nums)
	if len1 == 1 {
		return nums[0]
	}

	return find(nums, 0, len1-1, len1, k)
}

func find(nums []int, start int, end int, len1 int, target int) int {
	pivot := partition(nums, start, end)
	temp := pivot + target
	if temp == len1 {
		return nums[pivot]
	}
	if temp > len1 {
		return find(nums, start, pivot-1, len1, target)
	}

	return find(nums, pivot+1, end, len1, target)
}

func partition(nums []int, start int, end int) int {
	temp := nums[end]
	fast, slow := start, start
	for fast < end {
		if nums[fast] < temp {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}

	nums[slow], nums[end] = nums[end], nums[slow]
	return slow
}

func partition1(nums []int, lo, hi int) int {
	i := lo + 1
	j := hi
	for {
		for i < hi && less(nums[i], nums[lo]) {
			i++
		}
		for j > lo && less(nums[lo], nums[j]) {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}

func less(v, w int) bool {
	return v < w
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	//nums := []int{3, 2, 1, 5, 6, 4}
	//fmt.Println(findKthLargest(nums, 4))
	k := 4
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		j := partition1(nums, lo, hi)

		if j < k {
			lo = j + 1
		} else if j > k {
			hi = j - 1
		} else {
			break
		}
	}

	fmt.Println(nums[len(nums)-k])
}
