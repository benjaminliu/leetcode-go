package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	flag := 1
	if x < 0 {
		flag = -1
		x = -x
	}
	res := 0

	for x > 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res > math.MaxInt32 {
		return 0
	}
	return int(flag * res)
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}
