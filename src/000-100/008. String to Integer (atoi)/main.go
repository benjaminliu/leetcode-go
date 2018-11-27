package main

import (
	"math"
	"strconv"
)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	length := len(str)
	idx := 0
	for idx < length && str[idx] == 32 {
		idx ++
	}
	if idx+1 < length && (str[idx] == '+' || str[idx] == '-') {
		if str[idx+1] == '+' || str[idx+1] == '-' {
			return 0
		}
	}

	flag := 1
	if idx < length && str[idx] == '+' {
		idx++
	}
	if idx < length && str[idx] == '-' {
		flag = -1
		idx ++
	}

	sum := 0
	for idx < length && str[idx] >= 48 && str[idx] <= 57 {
		n, _ := strconv.Atoi(string(str[idx]))
		sum = sum*10 + n
		idx ++
		if sum > math.MaxInt32+1 {
			break
		}
	}

	res := flag * sum
	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}

func main() {

}
