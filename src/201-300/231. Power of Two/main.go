package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	if n < 3 {
		return true
	}
	if n%2 != 0 {
		return false
	}

	for n%2 == 0 {
		x := 2
		for n%x == 0 {
			n = n / x
			if n < x {
				break
			}
			x *= x
		}
	}
	if n > 1 {
		return false
	}
	return true
}

// if n is power of 2,  only the leftest digit is 1, others are 0
// n - 1 is all 1
//n & (n-1)  is 0
func isPowerOfTwo1(n int) bool {
	if n < 1 {
		return false
	}

	return n&(n-1) == 0
}

func main() {
	fmt.Println(isPowerOfTwo(24))
}
