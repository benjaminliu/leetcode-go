package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	if denominator == 0 {
		return ""
	}

	flag := 1
	var up, down int64
	if numerator < 0 {
		flag *= -1
		up = -int64(numerator)
	} else {
		up = int64(numerator)
	}

	if denominator < 0 {
		flag *= -1
		down = -int64(denominator)
	} else {
		down = int64(denominator)
	}

	res := ""
	if flag < 0 {
		res += "-"
	}

	//integral part
	val := up / down
	res += strconv.FormatInt(val, 10)


	// integral part,
	//up mode down,  and remainder * 10,  so we only need to process integral part
	up = (up % down) * 10
	if up ==0 {
		return res
	}
	res += "."

	maps := make(map[int64]int)

	beginIdx := -1
	idx := 0

	nums := make([]int64, 0)

	for up != 0 {
		if value, ok := maps[up]; ok {
			beginIdx = value
			break
		} else {
			maps[up] = idx
			idx++
			val = up / down
			up = (up % down) * 10
			nums = append(nums, val)
		}
	}
	for i := 0; i < idx; i++ {
		if i == beginIdx {
			res += "("
		}

		res += strconv.FormatInt(nums[i], 10)
	}
	if beginIdx != -1 {
		res += ")"
	}
	return res
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
}
