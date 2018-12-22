package main

import "fmt"

func isOneEditDistance(s, t string) bool {
	len1, len2 := len(s), len(t)
	diff := len1 - len2
	if diff > 1 || diff < -1 {
		return false
	}
	if diff == 0 {
		temp := 0
		for i := 0; i < len1; i++ {
			if s[i] != t[i] {
				temp ++
				if temp > 1 {
					return false
				}
			}
		}
		return true
	} else {
		if len1 > len2 {
			return isOneMore(s, t)
		} else {
			return isOneMore(t, s)
		}
	}
}
func isOneMore(longger string, shorter string) bool {
	len1 := len(shorter)
	idxL, idxS := 0, 0
	removed := false
	for idxS < len1 {
		if longger[idxL] == shorter[idxS] {
			idxL++
			idxS++
			continue
		}
		if removed == true {
			return false
		}
		removed = true
		idxL ++
	}
	return true
}

func main() {
	fmt.Println(isOneEditDistance("ab", "acb"))
	fmt.Println(isOneEditDistance("ad", "acb"))
	fmt.Println(isOneEditDistance("1203", "1213"))
}
