package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {

	flag := 1
	if dividend < 0 {
		flag *= -1
		dividend *= -1
	}
	if divisor < 0 {
		flag *= -1
		divisor *= -1
	}

	if dividend < divisor {
		return 0
	}
	if dividend == divisor {
		return flag
	}

	temp := 1
	temp1 := divisor
	for temp1 < dividend {
		temp1 = temp1 << 1
		temp = temp << 1
	}
	temp = temp >> 1
	temp1 = temp1 >> 1
	res := 0
	for dividend > 0 {
		dividend = dividend - temp1
		res += temp

		for temp > 1 && dividend < temp1 {
			temp1 = temp1 >> 1
			temp = temp >> 1
		}
		if dividend < temp1 {
			break
		}
	}
	result := res * flag
	if result > math.MaxInt32{
		return math.MaxInt32
	}
	if result < math.MinInt32{
		return math.MinInt32
	}
	return result
}

func main() {
	fmt.Println(divide(7, -3))
}
