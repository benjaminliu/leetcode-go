package main

func canJump(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return true
	}

	max := nums[0]
	if max == 0 {
		return false
	}
	
	last := length - 1
	for i := 1; i < length; i++ {
		temp := max - 1
		if temp > nums[i] {
			max = temp
		} else {
			max = nums[i]
		}
		if i != last {
			if max == 0 {
				return false
			}
		}
	}
	return true
}

func main() {

}
