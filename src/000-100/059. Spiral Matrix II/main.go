package main

import "fmt"

func generateMatrix(n int) [][]int {
	if n == 0 {
		return make([][]int, 0)
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	if n == 1 {
		res[0][0] = 1
		return res
	}

	i, j := 0, 0
	direction := 0
	left, right, top, down := 0, n-1, 0, n-1
	end := n * n
	idx := 1
	for idx <= end {
		res[i][j] = idx
		idx ++
		if direction == 0 {
			//go right
			if j < right {
				j++
			} else {
				i++
				top++
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
			if i > top {
				i--
			} else {
				j++
				left++
				changDirection(&direction)
			}
		}
	}
	return res
}

func changDirection(direction *int) {
	*direction++
	*direction %= 4
}

func main() {
	fmt.Println(generateMatrix(3))
}
