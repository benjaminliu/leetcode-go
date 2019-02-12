package main

import "fmt"

func firstBadVersion(n int) int {
	if isBadVersion(1) {
		return 1
	}
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)>>1
		if isBadVersion(mid) == false {
			if mid == right {
				return -1
			}

			if isBadVersion(mid + 1) {
				return mid + 1
			}

			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

func firstBadVersion1(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)>>1
		if isBadVersion(mid) == true {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	if version >= 5 {
		return true
	}
	return false
}

func main() {
	fmt.Println(firstBadVersion(5))
}
