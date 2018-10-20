package main

import (
	"bytes"
	"fmt"
)

var keys = []string{"MMM", "MM", "M",
	"CM", "DCCC", "DCC", "DC", "D", "CD", "CCC", "CC", "C",
	"XC", "LXXX", "LXX", "LX", "L", "XL", "XXX", "XX", "X",
	"IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}

var values = []int{3000, 2000, 1000,
	900, 800, 700, 600, 500, 400, 300, 200, 100,
	90, 80, 70, 60, 50, 40, 30, 20, 10,
	9, 8, 7, 6, 5, 4, 3, 2, 1}

func intToRoman(num int) string {
	length := len(keys)
	var buf bytes.Buffer
	for i := 0; i < length; i++ {
		if num >= values[i] {
			buf.WriteString(keys[i])
			num -= values[i]
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(intToRoman(1994))
}
