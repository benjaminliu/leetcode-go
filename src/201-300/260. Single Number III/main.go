package main

import "fmt"

func singleNumber(nums []int) []int {
	diff := 0
	for _, value := range nums {
		diff ^= value
	}

	diff &= -diff

	res := []int{0, 0}

	for _, value := range nums {
		if value&diff == 0 {
			res[0] ^= value
		} else {
			res[1] ^= value
		}
	}

	return res
}

func main() {
	temp := 3 ^ 5
	fmt.Println(temp)
	
}
