package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	len1 := len(s)
	len2 := len(t)
	if len1 != len2 {
		return false
	}

	if len1 < 2 {
		return true
	}

	maps1 := make(map[byte]int)
	maps2 := make(map[byte]int)

	maps1[s[0]] = 0
	maps2[t[0]] = 0
	for i := 1; i < len1; i++ {
		temp1 := s[i]
		temp2 := t[i]

		idx1, ok1 := maps1[temp1]
		idx2, ok2 := maps2[temp2]
		if ok1 != ok2 {
			return false
		}

		if idx1 != idx2 {
			return false
		}
		maps1[temp1] = i
		maps2[temp2] = i
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("ab", "aa"))
	fmt.Println(isIsomorphic("egg", "add"))
}
