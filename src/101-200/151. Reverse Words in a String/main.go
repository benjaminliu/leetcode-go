package main

import "fmt"

func reverseWords(s string) string {
	if len(s) < 2 {
		return s
	}

	len1 := len(s)
	cs := []byte(s)

	reverse(cs, 0, len1-1)

	start := 0
	for i := 1; i < len1; i++ {
		if cs[i] == ' ' {
			end := i-1
			reverse(cs, start, end)
			start = end + 2
		}
	}
	end := len1-1
	reverse(cs,start,end)
	return string(cs)
}

func reverse(cs []byte, left int, right int) {
	if left >= right {
		return
	}

	for left < right {
		cs[left],cs[right] = cs[right],cs[left]
		left ++
		right--
	}
}

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}
