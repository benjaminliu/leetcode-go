package main

import "fmt"

func numIslands(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}
	count := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] =='1' {
				count++
				dfs(grid, i,j, row, col)
			}
		}
	}

	return count
}
func dfs(grid [][]byte, i int, j int, row int, col int) {
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = 2

	if i > 0 {
		dfs(grid, i-1, j, row, col)
	}
	if i < row-1 {
		dfs(grid, i+1, j, row, col)
	}
	if j > 0 {
		dfs(grid, i, j-1, row, col)
	}
	if j < col-1 {
		dfs(grid, i, j+1, row, col)
	}
}

func main() {
	//grid := [][]byte{
	//	{1, 1, 0, 0, 0},
	//	{1, 1, 0, 0, 0},
	//	{0, 0, 1, 0, 0},
	//	{0, 0, 0, 1, 1},
	//}

	//grid := [][]byte{
	//	{'1', '1', '1', '1', '0'},
	//	{'1', '1', '0', '1', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '0', '0', '0'},
	//}

	grid := [][]byte{
		{'1'},
		{'0'},
		{'1'},
		{'0'},
		{'1'},
		{'1'},
	}

	fmt.Println(numIslands(grid))
}
