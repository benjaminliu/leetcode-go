package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	if rowIndex == 0 {
		return []int{1}
	}

	group := make([]int, rowIndex+1)
	group[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			group[j] += group[j-1]
		}
	}
	return group
}

func main() {
	fmt.Println(getRow(4))
}
