package main

import "fmt"

func restoreIpAddresses(s string) []string {
	len1 := len(s)
	res := make([]string, 0)
	if len1 > 12 || len1 < 4 {
		return res
	}

	group := make([]byte, len1+3)
	backtracking(s, 0, 0, 4, group, &res)

	return res
}
func backtracking(s string, idxS int, idxG int, count int, group []byte, res *[]string) {
	len1 := len(s)
	if idxS == len1 {
		if count == 0 {
			*res = append(*res, string(group))
			return
		} else {
			return
		}
	}
	left := len1 - idxS
	if left < count || left > count*3 {
		return
	}

	sum := 0
	for i := 0; i < 3; i++ {
		curS := idxS + i
		if curS >= len1 {
			break
		}
		curG := idxG + i
		group[curG] = s[curS]
		sum = sum*10 + c2i(s[curS])
		if sum > 255 {
			break
		}
		if count > 1 {
			group[curG+1] = '.'
			backtracking(s, curS+1, curG+2, count-1, group, res)
		} else {
			backtracking(s, curS+1, curG+1, count-1, group, res)
		}
		if sum == 0 {
			break
		}
	}
}

func c2i(c byte) int {
	return int(c - '0')
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
