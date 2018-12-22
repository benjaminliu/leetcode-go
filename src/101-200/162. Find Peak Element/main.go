package main

func findPeakElement(nums []int) int {
	last := len(nums) - 1
	if last == 0 {
		//len == 1
		return 0
	}

	if nums[0] > nums[1] {
		return 0
	}

	if nums[last] > nums[last-1]{
		return last
	}

	for i := 1; i < last; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}

func main() {

}
