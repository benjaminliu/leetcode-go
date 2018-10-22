package main

import (
	"fmt"
)

var maps = map[uint8][]rune{
	'2': []rune{'a', 'b', 'c'},
	'3': []rune{'d', 'e', 'f'},
	'4': []rune{'g', 'h', 'i'},
	'5': []rune{'j', 'k', 'l'},
	'6': []rune{'m', 'n', 'o'},
	'7': []rune{'p', 'q', 'r', 's'},
	'8': []rune{'t', 'u', 'v'},
	'9': []rune{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	length := len(digits)
	if length == 0 {
		return []string{}
	}

	count := 1
	for i := 0; i < length; i++ {
		count *= len(maps[digits[i]])
	}

	res := make([]string, count)
	no := 0
	idx := 0
	buf := make([]rune, length)
	helper(digits, buf, idx, length, res, &no)

	return res
}

func helper(digits string, buf []rune, idx int, length int, res []string, no *int) {
	if idx == length {
		res[*no] = string(buf)
		*no++
		return
	}

	items := maps[digits[idx]]
	for i, l := 0, len(items); i < l; i++ {
		buf[idx] = items[i]
		helper(digits, buf, idx+1, length, res, no)
	}
}

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}
