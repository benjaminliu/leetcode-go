package main

import "fmt"

func isPalindrome(s string) bool {
	len1 := len(s)
	left ,right :=0,len1-1

	for left < right {
		leftNo := toNum(s[left])
		for left < right && leftNo == -1{
			left++
			leftNo  = toNum(s[left])
		}

		rightNo := toNum(s[right])
		for left < right && rightNo == -1{
			right--
			rightNo = toNum(s[right])
		}
		if left == right{
			return true
		}
		if leftNo != rightNo{
			return false
		}
		left++
		right--
	}
	return true
}

func toNum(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}

	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 10)
	}

	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 10)
	}
	return -1
}

func main() {
fmt.Println(isPalindrome("0P"))
}
