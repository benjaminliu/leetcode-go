package main

func singleNumber(nums []int) int {
	len1 := len(nums)

	cur := nums[0]
	for i:=1;i< len1;i++{
		cur ^= nums[i]
	}
	return cur
}

func main() {
	
}
