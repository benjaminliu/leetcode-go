package main

import "fmt"

func climbStairs(n int) int {
	if n <2 {
		return n
	}
	if n == 2 {
		return 2
	}

	prepre :=1
	pre :=2
	cur :=0
	for i:=2;i<n;i++{
		cur = prepre +pre
		prepre = pre
		pre = cur
	}
	return cur
}

func main() {
	fmt.Println(climbStairs(4))
}
