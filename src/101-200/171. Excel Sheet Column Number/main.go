package main

import "fmt"

func titleToNumber(s string) int {
	len1 :=len(s)
	if len1 ==0 {
		return 0
	}
	res :=0
	for _,value :=range s{
		res = res * 26 + int(value - 'A') + 1
	}
	return res
}

func main() {
	fmt.Println(titleToNumber("A"))

	fmt.Println(titleToNumber("AB"))
}
