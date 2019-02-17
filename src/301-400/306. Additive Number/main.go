package main

import (
	"fmt"
)

func isAdditiveNumber(num string) bool {
	n := len(num)
	if n < 3 {
		return false
	}

	for i, end := 1, (n-1)/2; i <= end; i++ {
		//if first num has 2 or more digit, the first digit cannot be 0, here we return because the first digit of first num is the same
		if num[0] == '0' && i > 1 {
			return false
		}

		for j := i + 1; (n-j) >= i && (n-j) >= (j-i); j++ {
			// if second num has 2 or more digit, the first digit cannot be 0, here we continue because the first digit of second num changes ( start from 1, or start form 2 or ...)
			if num[i] == '0' && (j-i) > 1 {
				continue
			}

			pre := parseInt(num, 0, i)
			cur := parseInt(num, i, j)

			//after we choose the first and second num, we just validate the rule: third = first + second
			if isAdditive(num, n, j, pre, cur) {
				return true
			}
		}
	}

	return false
}

func parseInt(num string, start, end int) int {
	temp := 0

	for i := start; i < end; i++ {
		temp = temp*10 + byte2int(num[i])
	}
	return temp
}

func isAdditive(num string, n, idx, pre, cur int) bool {
	sum := pre + cur

	for idx < n {
		next := byte2int(num[idx])
		// if third num has 2 or more digit, the first digit cannot be 0
		if next == 0 && sum > 0 {
			return false
		}
		for next < sum {
			idx++
			if idx == n {
				return false
			}
			next = next*10 + byte2int(num[idx])
		}
		if next != sum {
			return false
		}
		idx ++
		pre = cur
		cur = sum
		sum = pre + cur
	}
	return true
}

func byte2int(b byte) int {
	return int(b - '0')
}

func main() {
	//fmt.Println(isAdditiveNumber("112358"))
	//fmt.Println(isAdditiveNumber("199100199"))
	//fmt.Println(isAdditiveNumber("1203"))
	//fmt.Println(isAdditiveNumber("121224036"))
	fmt.Println(isAdditiveNumber("198019823962"))
}
