package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	len := len(s)
	str := []byte(s)
	idx := 0
	maxLen, start := 1, 0
	for idx < len {
		if maxLen/2 >= len-idx {
			break
		}
		j, k := idx, idx
		for k < len-1 && str[k+1] == str[k] {
			k++
		}
		idx = k + 1
		for k < len-1 && j > 0 && str[j-1] == str[k+1] {
			j--
			k++
		}
		if k-j+1 > maxLen {
			start = j
			maxLen = k - j + 1
		}
	}
	return string(str[start : start+maxLen])
}

func main() {
	s := longestPalindrome("babad")
	fmt.Println(s)
}
