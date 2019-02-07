package main

import (
	"fmt"
	"strconv"
)

var maps map[string][]int = make(map[string][]int)

func diffWaysToCompute(input string) []int {
	if value, ok := maps[input]; ok {
		return value
	}

	len1 := len(input)

	res := make([]int, 0)
	for i := 0; i < len1; i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			left := diffWaysToCompute(string(input[0:i]))
			right := diffWaysToCompute(string(input[i+1:]))

			for _, leftValue := range left {
				for _, rightValue := range right {
					if input[i] == '+' {
						res = append(res, leftValue+rightValue)
					}
					if input[i] == '-' {
						res = append(res, leftValue-rightValue)
					}
					if input[i] == '*' {
						res = append(res, leftValue*rightValue)
					}
				}
			}
		}
	}

	if len(res) == 0 {
		value, _ := strconv.Atoi(input)
		res = append(res, value)
	}

	maps[input] = res

	return res
}

func main() {
	input := "2*3-4*5"
	fmt.Println(diffWaysToCompute(input))
}
