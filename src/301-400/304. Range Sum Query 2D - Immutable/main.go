package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
	row    int
	col    int
}

func Constructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	if row == 0 {
		return NumMatrix{matrix: nil, row: 0, col: 0}
	}

	col := len(matrix[0])
	if col == 0 {
		return NumMatrix{matrix: nil, row: 0, col: 0}
	}

	temp := make([][]int, row)
	for i := 0; i < row; i++ {
		temp[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		temp[i][0] = matrix[i][0]
	}

	for i := 0; i < row; i++ {
		for j := 1; j < col; j++ {
			temp[i][j] = temp[i][j-1] + matrix[i][j]
		}
	}

	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			temp[i][j] += temp[i-1][j]
		}
	}

	return NumMatrix{matrix: temp, row: row, col: col}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 < 0 || col1 < 0 || row2 > this.row || col2 > this.col || this.row == 0 || this.col == 0 {
		return 0
	}

	if row1 == 0 && col1 == 0 {
		return this.matrix[row2][col2]
	}

	if row1 == 0 {
		return this.matrix[row2][col2] - this.matrix[row2][col1-1]
	}

	if col1 == 0 {
		return this.matrix[row2][col2] - this.matrix[row1-1][col2]
	}

	return this.matrix[row2][col2] - this.matrix[row2][col1-1] - this.matrix[row1-1][col2] + this.matrix[row1-1][col1-1]
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	obj := Constructor(matrix)
	fmt.Println(obj.SumRegion(2, 1, 4, 3))
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
	fmt.Println(obj.SumRegion(1, 2, 2, 4))
}
