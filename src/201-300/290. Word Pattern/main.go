package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	len1 := len(pattern)
	if len1 != len(strs) {
		return false
	}

	mapsPattern := make(map[byte]int)
	mapsPattern[pattern[0]] = 0

	mapsStr := make(map[string]int)
	mapsStr[strs[0]] = 0

	for i := 1; i < len1; i++ {
		v1, ok1 := mapsPattern[pattern[i]]
		v2, ok2 := mapsStr[strs[i]]

		if ok1 != ok2 {
			return false
		}
		if ok1 {
			if v1 != v2 {
				return false
			}
		} else {
			mapsPattern[pattern[i]] = i
			mapsStr[strs[i]] = i
		}
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
}
