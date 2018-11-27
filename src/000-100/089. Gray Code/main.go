package main

import "fmt"

func grayCode(n int) []int {
	res := make([]int, 0)
	res = append(res, 0)
	if n == 0 {
		return res
	}
	res = append(res, 1)

	num := 1 << uint(n)
	temp := 0
	add := 0
	for i := 2; i < num; i++ {
		temp = twoSquareRoot(i)
		id := 1<<uint(temp) - 1 - i
		add = res[id] + 1<<uint(temp-1)
		res = append(res, add)
	}
	return res
}

func twoSquareRoot(n int) int {
	res := 0
	for n > 0 {
		n = n / 2
		res++
	}
	return res
}

func grayCode1(n int) []int {
	if n == 0 {
		return []int{0}
	}
	g := grayCode(n-1)
	l := len(g)
	for i:=0;i<l;i++ {
		g = append(g, int(uint(1)<<uint(n-1)) + g[l-i-1])
	}
	return g
}

func main() {
	fmt.Println(grayCode(2))
}
