package main

import "fmt"

func multiply(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)

	if len1 < len2 {
		return multiply(num2, num1)
	}
	if num1 == "0" {
		return "0"
	}

	if num2 == "0" {
		return "0"
	}

	totalLen := len1 + len2

	res := make([]int, totalLen)

	idx := totalLen - 1
	idx2 := len2 - 1
	for idx2 >= 0 {
		if num2[idx2] == '0' {
			idx --
			idx2--
		} else {
			break
		}
	}
	increase := 0
	idxTemp := idx
	n2 := int(num2[idx2] - 48)
	for idx1 := len1 - 1; idx1 >= 0; idx1-- {
		temp := int(num1[idx1]-48)*n2 + increase
		res[idxTemp] = temp % 10
		increase = temp / 10
		idxTemp--
	}
	res[idxTemp] = increase
	idx2--
	idx--
	for idx2 >= 0 {
		idxTemp = idx
		increase = 0
		n2 := int(num2[idx2] - 48)
		for idx1 := len1 - 1; idx1 >= 0; idx1-- {
			cur := int(num1[idx1]-48) * n2
			temp := cur + increase + res[idxTemp]
			res[idxTemp] = temp % 10
			increase = temp / 10
			idxTemp--
		}
		res[idxTemp] = increase
		idx2 --
		idx--
	}
	temp := make([]byte, 0)
	i := 0
	for ; i < totalLen; i++ {
		if res[i] == 0 {

		} else {
			break
		}
	}

	for ; i < totalLen; i++ {
		temp = append(temp, byte(res[i]+48))
	}

	return string(temp)
}

func multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1 := len(num1)
	len2 := len(num2)

	res := make([]int, len1+len2)

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			res[i+j+1] += int(num1[i]-48) * int(num2[j]-48)
		}
	}

	for i := len(res) - 1; i > 0; i-- {
		if res[i] >= 10 {
			res[i-1] += res[i] / 10
			res[i ] %= 10
		}
	}

	temp := make([]byte, 0)
	i := 0
	//only the first could be 0
	if res[i] == 0 {
		i++
	}

	for ; i < len(res); i++ {
		temp = append(temp, byte(res[i]+48))
	}

	return string(temp)
}

/*
整体思路，个位数相乘，但是先不用进位，保存在int 数组里
*/
func multiply1(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	tmp := make([]int, len(num1)+len(num2), len(num1)+len(num2))

	for i := 0; i < len(num1); i ++ {
		for j := 0; j < len(num2); j++ {
			tmp[i+j+1] = tmp[i+j+1] + int(num1[i]-'0')*int(num2[j]-'0')
		}
	}

	mod := 0
	i := len(tmp) - 1

	for ; i >= 0; i-- {
		c := tmp[i]
		tmp[i] = (c + mod) % 10
		mod = (c + mod) / 10
	}

	if tmp[0] == 0 {
		tmp = tmp[1:len(tmp)]
	}
	newS := make([]byte, 0, len(tmp))
	for _, v := range tmp {
		newS = append(newS, byte(v)+'0')
	}
	return string(newS)
}

func main() {
	res := multiply2("5", "12")
	fmt.Println(res)
}
