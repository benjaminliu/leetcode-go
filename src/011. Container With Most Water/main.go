package main

import "fmt"

func minInt(left int, right int) int {
	if left > right {
		return right
	}
	return left
}

func maxArea(height []int) int {
	length := len(height)
	if length < 2 {
		return 0
	}
	leftIdx := 0
	rightIdx := length - 1
	lastLeft := height[leftIdx]
	lastRight := height[rightIdx]
	max := minInt(lastLeft, lastRight) * (rightIdx - leftIdx)
	for leftIdx < rightIdx {
		if lastLeft > lastRight {
			rightIdx--
			lastRight = height[rightIdx]
		} else {
			leftIdx++
			lastLeft = height[leftIdx]
		}
		temp := minInt(lastLeft, lastRight) * (rightIdx - leftIdx)
		if temp > max {
			max = temp
		}
	}
	return max
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
