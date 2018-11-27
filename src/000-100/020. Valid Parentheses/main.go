package main

import "fmt"

func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}
	if length%2 == 1 {
		return false
	}

	stack := []byte{}
	for i := 0; i < length; i++ {
		cur := s[i]
		if cur == '(' || cur == '{' || cur == '[' {
			stack = append(stack, cur)
		} else {
			if len(stack ) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if cur == ')' && top != '(' {
				return false
			}
			if cur == '}' && top != '{' {
				return false
			}
			if cur == ']' && top != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
