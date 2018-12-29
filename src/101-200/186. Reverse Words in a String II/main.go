package main

import "fmt"

func reverseWords(ss string) string {
	len1 := len(ss)
	if len1 < 2 {
		return ss
	}

	s := []byte(ss)
	start := 0
	for i := 0; i < len1; i++ {
		if s[i] == ' ' {
			reverse(s, start, i-1)
			start = i + 1
		}
	}
	reverse(s, start, len1-1)

	reverse(s, 0, len1-1)
	return string(s)
}
func reverse(s []byte, start int, end int) {

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func main() {
	s := "the sky is blue";
	s = reverseWords(s)
	fmt.Println(s)
}
