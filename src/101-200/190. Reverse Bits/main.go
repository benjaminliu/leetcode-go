package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var res uint32
	idx := 32
	for idx > 0 {
		temp := num % 2
		num >>= 1
		res = res<<1 + temp
		idx--
	}
	return res
}

func main() {
	var num uint32 = uint32(43261596)
	fmt.Println(reverseBits(num))
}
