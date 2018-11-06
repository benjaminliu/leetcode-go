package main

func minPathSum(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}

	for i := 1; i < row; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for i := 1; i < col; i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if grid[i-1][j] > grid[i][j-1] {
				grid[i] [j] += grid[i][j-1]
			} else {
				grid[i] [j] += grid[i-1][j]
			}
		}
	}

	return grid[row-1][col-1]
}

func main() {

}
