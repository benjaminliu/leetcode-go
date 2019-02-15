package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	if triangle == nil || len(triangle) == 0 {
		return 0
	}

	len1 := len(triangle)

	for i := len1 - 1; i > 0; i-- {
		cur := triangle[i]
		pre := triangle[i-1]
		for j := 1; j <= i; j++ {
			pre[j-1] += minInt(cur[j-1], cur[j])
		}
	}
	return triangle[0][0]
}

//dp
func minimumTotal1(triangle [][]int) int {
	len1 := len(triangle)
	dp := triangle[len1-1]

	for layer := len1-2;layer>=0;layer--{
		for i:=0;i<=layer;i++{
			dp[i] = minInt(dp[i],dp[i+1]) + triangle[layer][i]
		}
	}
	return dp[0]
}

func minInt(left int, right int) int {
	if left <= right {
		return left
	}
	return right
}



func main() {
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}

	fmt.Println(minimumTotal(triangle))
}
