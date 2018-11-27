package main

import "fmt"

func strStr(haystack string, needle string) int {
	lenN := len(needle)
	if lenN == 0 {
		return 0
	}

	lenH := len(haystack)
	if lenN > lenH {
		return -1
	}
	sumN := 0
	sumH := 0
	highest := 1
	for i := 0; i < lenN; i++ {
		highest *= 26
		sumN = sumN*26 + c2i(needle[i])
		sumH = sumH*26 + c2i(haystack[i])
	}
	highest /= 26

	idx := 0
	if sumH == sumN {
		return idx
	}
	for i := lenN; i < lenH; i++ {
		sumH -= c2i(haystack[idx]) * highest
		sumH = sumH*26 + c2i(haystack[i])
		idx++
		if sumH == sumN {
			return idx
		}
	}
	return -1
}

func c2i(c byte) int {
	return int(c) - 97
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}
