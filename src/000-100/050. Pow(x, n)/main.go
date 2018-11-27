package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	var flag float64 = 1
	if x < 0 {
		x = -x
		flag = -1
	}

	y := myPow(x, n/2)
	if n%2 == 0 {
		return y * y
	} else {
		return y * y * x * flag
	}
}

func main() {
	fmt.Println(myPow(2 , 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2, -2))
}
