package main

import "fmt"

func exist(board [][]byte, word string) bool {
	length := len(word)
	if length == 0 {
		return true
	}

	row := len(board)
	if row == 0 {
		return false
	}
	col := len(board[0])
	if col == 0 {
		return false
	}

	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				if dfs(board, visited, i, j, row, col, word, 0, length) {
					return true
				}
			}
		}
	}
	return false
}
func dfs(board [][]byte, visited [][]bool, i int, j int, row int, col int, word string, idx int, length int) bool {
	if idx == length {
		return true
	}

	if visited[i][j] == true {
		return false
	}

	if board[i][j] != word[idx] {
		return false
	}
	idx ++
	if idx == length {
		return true;
	}

	visited[i][j] = true

	if i > 0 {
		if dfs(board, visited, i-1, j, row, col, word, idx, length) {
			return true
		}
	}
	if i < row-1 {
		if dfs(board, visited, i+1, j, row, col, word, idx, length) {
			return true
		}
	}
	if j > 0 {
		if dfs(board, visited, i, j-1, row, col, word, idx, length) {
			return true
		}
	}
	if j < col-1 {
		if dfs(board, visited, i, j+1, row, col, word, idx, length) {
			return true
		}
	}

	visited[i][j] = false
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))
}
