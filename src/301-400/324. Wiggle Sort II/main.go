package main

import "sort"

func wiggleSort2(nums []int) {
	if len(nums) <= 1 {
		return
	}
	tmp := []int{}
	tmp = append(tmp, nums...)
	sort.Ints(tmp)

	n := len(nums)
	left, right := (n+1)/2, n

	for i := range nums {
		if i&1 == 0 {
			left--
			nums[i] = tmp[left]
		} else {
			right--
			nums[i] = tmp[right]
		}
	}
}

func wiggleSort1(nums []int) {
	if len(nums) <= 1 {
		return
	}
	n := len(nums)
	sort.Ints(nums)
	mid := n/2 + n%2
	var res []int

	for i := 0; i < mid; i++ {
		res = append(res, nums[mid-i-1])
		if mid+i < n {
			res = append(res, nums[n-i-1])
		}
	}
	for i := 0; i < len(res); i++ {
		nums[i] = res[i]
	}
}

func wiggleSort(nums []int) {
	n := len(nums)
	mid := findKthLargest(nums, (n+1)/2)

	left, i, right := 0, 0, n-1

	for i <= right {
		temp := newIndex(i, n)
		if nums[temp] > mid {
			swap(nums, newIndex(left, n), newIndex(i, n))
			left++
			i++
		} else if nums[temp] < mid {
			swap(nums, newIndex(right, n), newIndex(i, n))
			right--
		} else {
			i++
		}
	}
}
func swap(nums []int, a int, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func newIndex(idx, n int) int {
	return (1 + 2*idx) % (n | 1)
}

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

func main() {

}
