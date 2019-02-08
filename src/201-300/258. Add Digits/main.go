package main

import "fmt"

func addDigits(num int) int {
	temp := 0
	for num > 10 {
		for num > 10 {
			temp += num % 10
			num /= 10
		}
		num += temp
		temp = 0
	}
	return num
}
func addDigits1(num int) int {
	return 1 + (num-1)%9
}

func main() {
	fmt.Println(addDigits(38))
}
