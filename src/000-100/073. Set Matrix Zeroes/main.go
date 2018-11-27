package main

import "fmt"

func setZeroes(matrix [][]int) {
	row := len(matrix)
	if row == 0 {
		return
	}
	col := len(matrix[0])
	if col == 0 {
		return
	}

	visited := make([][]int, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if (visited[i][j] == 1) {
				continue
			}
			visited[i][j] = 1

			if matrix[i][j] == 0 {
				for x := 0; x < row; x++ {
					if matrix[x][j] != 0 {
						matrix[x][j] = 0
						visited[x][j] = 1
					}
				}
				for x := 0; x < col; x++ {
					if matrix[i][x] != 0 {
						matrix[i][x] = 0
						visited[i][x] = 1
					}
				}
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}
