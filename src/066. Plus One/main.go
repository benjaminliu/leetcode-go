package main

import "fmt"

func plusOne(digits []int) []int {
	length := len(digits)

	temp := digits[length-1] + 1
	increase := temp / 10
	digits[length-1] = temp % 10
	for i := length - 2; i >= 0; i-- {
		temp = digits[i] + increase
		increase = temp / 10
		digits[i] = temp % 10
	}
	if increase == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(plusOne(nums))
}
