package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n < 3 {
		return false
	}

	for n%3 == 0 {
		divide := 3
		for n%divide == 0 {
			n /= divide
			divide *= 3
		}
	}
	if n == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPowerOfThree(45))
}
