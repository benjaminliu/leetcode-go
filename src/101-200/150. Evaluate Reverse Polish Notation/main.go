package main

import "strconv"

func evalRPN(tokens []string) int {
	if tokens == nil || len(tokens) == 0 {
		return 0
	}

	stack := make([]int, 0)

	for _, value := range tokens {
		if value == "+" {
			right := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			left := stack[len(stack)-1]
			temp := left + right
			stack[len(stack)-1] = temp
		} else if value == "-" {
			right := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			left := stack[len(stack)-1]
			temp := left - right
			stack[len(stack)-1] = temp
		} else if value == "*" {
			right := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			left := stack[len(stack)-1]
			temp := left * right
			stack[len(stack)-1] = temp
		} else if value == "/" {
			right := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			left := stack[len(stack)-1]
			temp := left / right
			stack[len(stack)-1] = temp
		} else {
			intValue, _ := strconv.Atoi(value)
			stack = append(stack, intValue)
		}
	}
	return stack[0]
}

func main() {

}
