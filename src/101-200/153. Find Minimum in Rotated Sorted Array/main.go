package main

import "fmt"

func findMin(nums []int) int {
	len1 := len(nums)
	if len1 == 1 {
		return nums[0]
	}

	end := len1 - 1
	if nums[0] < nums[end] {
		return nums[0]
	}
	if nums[end-1] > nums[end] {
		return nums[end]
	}

	start := 0
	for start <= end {
		mid := start + (end-start)/2
		if mid> 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[start] < nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return 0
}

func findMin1(nums []int) int {
	start,end:=0,len(nums)-1
	for start < end {
		mid := start + (end - start)/2
		if nums[mid] > nums[end]{
			start = mid +1
		}else {
			end = mid
		}
	}
	return nums[start]
}

func main() {
	//nums := []int{4, 5, 6, 1, 2, 3}
	//nums := []int{5, 1, 2, 3, 4}
	//nums :=[]int{3,4,5,1,2}
	//nums :=[]int{4,5,6,7,0,1,2}
	//nums :=[]int{1,0}
	nums :=[]int{2,3,1}
	fmt.Println(findMin(nums))
}
