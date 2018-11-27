package main

import "fmt"

func addBinary(a string, b string) string {
	len1 := len(a)
	len2 := len(b)
	if len1 < len2 {
		return addBinary(b, a)
	}

	res := make([]byte, len1+1)
	idx := len1 - 1
	inc := 0
	for i := len2 - 1; i >= 0; i-- {
		temp := c2i(a[idx]) + c2i(b[i]) + inc
		res[idx+1] = i2c(temp % 2)
		inc = temp / 2
		idx--
	}

	for i := idx + 1; i > 0; i-- {
		temp := c2i(a[idx]) + inc
		res[idx+1] = i2c(temp % 2)
		inc = temp / 2
		idx --
	}
	if inc == 0 {
		return string(res[1:])
	} else {
		res[0] = i2c(1)
		return string(res)
	}
}

func c2i(c byte) int {
	return int(c - '0')
}

func i2c(i int) byte {
	return byte(i + '0')
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
