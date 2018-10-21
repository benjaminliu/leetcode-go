package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}

	first := strs[0]
	firstLen := len(first)
	idx := 0
out:
	for idx < firstLen {
		temp := first[idx]
		for i := 1; i < length; i++ {
			if idx == len(strs[i]) {
				break out
			}
			if temp != strs[i][idx] {
				break out
			}
		}
		idx ++
	}
	return first[0:idx]
}

func main() {
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{ "dog","racecar","car"}
	fmt.Println(longestCommonPrefix(strs))
}
