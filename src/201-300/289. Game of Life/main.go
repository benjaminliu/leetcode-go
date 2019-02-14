package main

import "fmt"

func gameOfLife(board [][]int) {
	row := len(board)
	col := len(board[0])

	if row == 0 || col == 0 {
		return
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			count := countOne(board, i, j, row, col)

			if count < 2 || count > 3 {
				//under-population or over-population
				if board[i][j] == 1 {
					board[i][j] = -1
				}
			} else {
				if count == 3 {
					//reproduction
					if board[i][j] == 0 {
						board[i][j] = 2
					}
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] > 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}

func countOne(board [][]int, i int, j int, row int, col int) int {
	count := 0

	if isOne(board, i-1, j-1, row, col) {
		count++
	}
	if isOne(board, i-1, j, row, col) {
		count++
	}
	if isOne(board, i-1, j+1, row, col) {
		count++
	}
	if isOne(board, i, j-1, row, col) {
		count++
	}
	if isOne(board, i, j+1, row, col) {
		count++
	}
	if isOne(board, i+1, j-1, row, col) {
		count++
	}
	if isOne(board, i+1, j, row, col) {
		count++
	}
	if isOne(board, i+1, j+1, row, col) {
		count++
	}
	return count
}

func isOne(board [][]int, i, j, row, col int) bool {
	if i < 0 || i == row || j < 0 || j == col {
		return false
	}

	if board[i][j] == 1 {
		return true
	}
	if board[i][j] == -1 {
		return true
	}
	return false
}

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	gameOfLife(board)
	fmt.Println(board)
}
