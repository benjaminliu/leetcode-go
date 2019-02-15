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

//dp
func climbStairs1(n int) int {
	maps := make(map[int]int)
	return helper(n,maps)
}
func helper(n int, maps map[int]int) int {
	if n<0{
		return 0
	}
	if n ==0 {
		return 1
	}

	if value,ok:=maps[n];ok{
		return value
	}

	value := helper(n-1,maps) + helper(n-2,maps)

	maps[n] =value
	return value
}

func main() {
	fmt.Println(climbStairs(4))
}
