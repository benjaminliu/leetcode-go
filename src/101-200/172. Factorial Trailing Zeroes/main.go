package main

import "fmt"

func trailingZeroes(n int) int {
	res := 0
	for n > 4 {
		n /= 5
		res += n
	}
	return res
}

func main() {
	fmt.Println(factorial(9))
	fmt.Println(factorial(10))
}

func factorial(n int) int64 {
	num := int64(n)
	var res int64 = 1
	for num > 1 {
		res *= num
		num--
	}
	return res
}
