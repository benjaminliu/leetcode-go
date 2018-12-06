package main

import "fmt"

func solve(board [][]byte) {
	row := len(board)
	if row == 0 {
		return
	}
	col := len(board[0])
	if col == 0 {
		return
	}
	visit := make([][]byte, row)
	for i := 0; i < row; i++ {
		visit[i] = make([]byte, col)
	}

	lastCol := col - 1

	for i := 0; i < row; i++ {
		dfs(board, i, 0, visit)
		dfs(board, i, lastCol, visit)
	}
	lastRow := row - 1
	for i := 0; i < col; i++ {
		dfs(board, 0, i, visit)
		dfs(board, lastRow, i, visit)
	}

	for i := 1; i < lastRow; i++ {
		for j := 1; j < lastCol; j++ {
			if visit[i][j] == 1 {
				continue
			}
			board[i][j] = 'X'
		}
	}
}
func dfs(board [][]byte, row int, col int, visit [][]byte) {
	if visit[row][col] == 1 {
		return
	}
	visit[row][col] = 1
	if board[row][col] == 'X'{
		return
	}
	if row > 0 {
		dfs(board, row-1, col, visit)
	}
	if row < len(board)-1 {
		dfs(board, row+1, col, visit)
	}
	if col > 0 {
		dfs(board, row, col-1, visit)
	}
	if col < len(board[0])-1 {
		dfs(board, row, col+1, visit)
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X','X'},
		{'X', 'O', 'O','X'},
		{'X', 'X', 'O','X'},
		{'X', 'O', 'X','X'},
	}

	solve(board)

	fmt.Println(board)
}
