package main

func moveZeroes(nums []int) {
	len1 := len(nums)
	zeroCount := 0
	i := 0
	for ; i < len1; i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			if zeroCount > 0 {
				nums[i-zeroCount] = nums[i]
			}
		}
	}
	for i = len1 - zeroCount; i < len1; i++ {
		nums[i] = 0
	}
}

func main() {

}
