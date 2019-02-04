package main

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	if col == 0 {
		return false
	}
	if matrix[0][0] > target {
		return false
	}
	if matrix[row-1][col-1] < target {
		return false
	}

	i := row - 1
	j := 0
	for i >= 0 && j < col {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			j ++
		} else {
			i--
		}
	}
	return false
}

func main() {

}
