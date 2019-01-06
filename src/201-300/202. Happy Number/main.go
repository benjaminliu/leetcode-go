package main

import "fmt"

func isHappy(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	maps := make(map[int]bool)
	target := n
	for target != 1 {
		temp := 0
		for target > 0 {
			x := target % 10
			temp += x * x
			target /= 10
		}
		if _, ok := maps[temp]; ok {
			return false
		}
		maps[temp] = true
		target = temp
	}
	return true
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}
