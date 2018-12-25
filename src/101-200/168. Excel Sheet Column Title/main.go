package main

import "fmt"

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}

	chars := make([]byte, 0)

	for n > 0 {
		n--
		temp := n % 26
		chars = append([]byte{byte(temp) + byte('A')}, chars...)
		n /= 26
	}

	return string(chars)
}

func main() {
	fmt.Println(convertToTitle(1))
}
