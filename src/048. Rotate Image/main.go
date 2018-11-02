package main

import "fmt"

func rotate(matrix [][]int) {
	length := len(matrix)
	if length < 2 {
		return
	}

	for row := 0; row < length; row ++ {
		for col := row + 1; col < length; col ++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	for row := 0; row < length; row ++ {
		left, right := 0, length-1
		for left < right {
			matrix[row][left], matrix[row] [right] = matrix[row] [right], matrix[row][left]
			left++
			right --
		}
	}
}

func main() {
	matrix := make([][]int, 3)
	matrix[0] = []int{1, 2, 3}
	matrix[1] = []int{4, 5, 6}
	matrix[2] = []int{7, 8, 9}
	rotate(matrix)
	fmt.Println(matrix)
}
