package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	if col == 0 {
		return false
	}
	if target < matrix[0][0] {
		return false
	}
	if target > matrix[row-1][col-1] {
		return false
	}

	count := row * col
	left, right := 0, count-1
	for left <= right {
		mid := left + (right-left)/2
		i := mid / col
		j := mid % col
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	matrix := [][]int {
		{1,3,5,7},
		{10,11,16,20},
		{23,30,34,50},
	}
	fmt.Println(searchMatrix(matrix,16))
}
