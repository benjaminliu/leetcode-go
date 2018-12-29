package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	len1 := len(s)
	if len1 < 11 {
		return nil
	}

	total := 0
	for i := 0; i < 10; i++ {
		total = total*10 + a2i(s[i])
	}
	maps := make(map[int]int)
	maps[total] = 1

	res := make([]string, 0)

	for i := 10; i < len1; i++ {
		temp := a2i(s[i-10]) * 1000000000
		total = total - temp
		total = total*10 + a2i(s[i])
		if value, ok := maps[total]; ok {
			if value == 1 {
				res = append(res, s[i-9:i+1])
			}
			maps[total] = value + 1
		} else {
			maps[total] = 1
		}
	}
	return res
}

func a2i(c byte) int {
	if c == 'A' {
		return 1
	}
	if c == 'C' {
		return 2
	}
	if c == 'G' {
		return 3
	}
	return 4
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))
}
