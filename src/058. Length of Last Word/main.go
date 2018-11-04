package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := len(s)

	idx := length - 1
	for idx >= 0 {
		if s[idx] == ' ' {
			idx --
		} else {
			break
		}
	}

	end := idx

	for idx >= 0 {
		if s[idx ] != ' ' {
			idx --
		} else {
			break
		}
	}

	start := idx
	if start == end {
		return 0
	}

	return end - start
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
