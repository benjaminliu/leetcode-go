package main

import "fmt"

func maxProduct(words []string) int {
	n := len(words)
	if n == 0 {
		return 0
	}

	list := make([]int, 0)

	for _, value := range words {
		list = append(list, string2int(value))
	}

	maxValue := 0
	last := n - 1
	for i := 0; i < last; i++ {
		for j := i + 1; j < n; j++ {
			if list[i]&list[j] == 0 {
				temp := len(words[i]) * len(words[j])
				if temp > maxValue {
					maxValue = temp
				}
			}
		}
	}
	return maxValue
}

func string2int(str string) int {
	res := 0
	n := len(str)
	if n == 0 {
		return res
	}

	for i := 0; i < n; i++ {
		res |= c2bit(str[i])
	}
	return res
}

func c2bit(c byte) int {
	shift := c - 'a'
	return int(1 << shift)
}

func main() {
	fmt.Println(string2int("abcd"))
	fmt.Println(string2int("abcde"))

	fmt.Println(c2bit('a'))
	fmt.Println(c2bit('b'))
	fmt.Println(c2bit('c'))
	fmt.Println(c2bit('d'))
	fmt.Println(c2bit('e'))
	fmt.Println(c2bit('z'))
}
