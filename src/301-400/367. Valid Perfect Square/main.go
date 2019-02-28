package main

import "fmt"

var squres []int = []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100, 121, 144, 169, 196, 225}
var base int = 15

func isPerfectSquare(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}

	if squres[base] > num {
		return binarySearch(squres, num)
	} else if squres[base] == num {
		return true
	}
	for squres[base] < num {
		base++
		squres = append(squres, base*base)
	}

	if squres[base] == num {
		return true
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	left, right := 2, len(arr)

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
}
