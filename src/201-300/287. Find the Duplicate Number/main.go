package main

func findDuplicate(nums []int) int {
	maps := make(map[int]int)

	for _, value := range nums {
		if no, ok := maps[value]; ok {
			return no
		} else {
			maps[value] = value
		}
	}
	return 0
}

func findDuplicate1(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}

func main() {

}
