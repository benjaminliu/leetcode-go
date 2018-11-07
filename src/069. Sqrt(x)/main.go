package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if 4 > x {
		return 1
	}

	end := x/2 + 1
	for i := 2; i <= end; i++ {
		square := i * i
		if square == x {
			return i
		}
		if square > x {
			return i - 1
		}
	}
	return 0
}

func mySqrt1(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}

	left, right := 2, x
	for left < right {
		if left*left <= x && (left+1)*(left+1) > x {
			return left
		}

		mid := left + (right-left)/2

		if mid*mid <= x {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(mySqrt1(8))
}
