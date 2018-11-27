package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	if row == 1 {
		return matrix[0]
	}
	col := len(matrix[0])
	if col == 0 {
		return []int{}
	}
	res := make([]int, 0)
	if col == 1 {
		for i := 0; i < row; i++ {
			res = append(res, matrix[i][0])
		}
		return res
	}

	direction := 0
	left, right, up, down := 0, col-1, 0, row-1
	totalLen := row * col
	idx := 0
	i, j := 0, 0
	for idx < totalLen {
		res = append(res, matrix[i][j])
		if direction == 0 {
			//go right
			if j < right {
				j++
			} else {
				i++
				up++
				changDirection(&direction)
			}
		} else if direction == 1 {
			//go down
			if i < down {
				i++
			} else {
				j--
				right--
				changDirection(&direction)
			}
		} else if direction == 2 {
			//go left
			if j > left {
				j--
			} else {
				i--
				down--
				changDirection(&direction)
			}
		} else {
			//go up
			if i > up {
				i--
			} else {
				j++
				left++
				changDirection(&direction)
			}
		}
		idx ++
	}

	return res
}

func changDirection(direction *int) {
	*direction++
	*direction %= 4
}

func main() {
	matrix := make([][]int, 3)
	matrix[0] = []int{1, 2, 3}
	matrix[1] = []int{4, 5, 6}
	matrix[2] = []int{7, 8, 9}
	fmt.Println(spiralOrder(matrix))
}
