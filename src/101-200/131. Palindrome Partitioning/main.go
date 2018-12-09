package main

import "fmt"

func partition(s string) [][]string {
	len1 := len(s)
	if len1 == 0 {
		return nil
	}

	res := make([][]string, 0)
	group := make([]string, 0)

	helper(s, 0, len1-1, &res, group)
	return res
}
func helper(s string, idx int, last int, res *[][]string, group []string) {
	if idx >  last {
		temp := make([]string, len(group))
		copy(temp, group)
		*res = append(*res, temp)
		return
	}
	for i := last; i >= idx; i-- {
		if isPalindrome(s, idx, i) {
			helper(s, i+1, last, res, append(group, s[idx:i+1]))
		}
	}
}

func isPalindrome(s string, left int, right int) bool {
	if right < left {
		return false
	}
	if right == left {
		return true
	}

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
