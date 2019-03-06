package main

//need more care
func kthSmallest(matrix [][]int, k int) int {
	row := len(matrix)
	col := len(matrix[0])
	lo := matrix[0][0]
	hi := matrix[row-1][col-1] + 1

	for lo < hi {
		mid := lo + (hi-lo)/2
		count := 0
		j := col - 1
		for i := 0; i < row; i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			count += (j + 1)
		}
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func kthSmallest1(matrix [][]int, k int) int {
	row := len(matrix)
	col := len(matrix[0])
	lo := matrix[0][0]
	hi := matrix[row-1][col-1] + 1

	for lo < hi {
		mid := lo + (hi-lo)/2
		count := 0
		for i := 0; i < row; i++ {
			temp := 0
			for j := 0; j < col; j ++ {
				if matrix[i][j] <= mid {
					temp++
				} else {
					break
				}
			}
			count += temp
		}
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {

}
