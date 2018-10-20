package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	temp := x % 10
	if temp == 0 {
		return false
	}

	y := 0
	for x > y {
		temp = x % 10
		x = x / 10
		if x == y {
			return true
		}
		y = y*10 + temp
		if x == y {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(11))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(112))
}
