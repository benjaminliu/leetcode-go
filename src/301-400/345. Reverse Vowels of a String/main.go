package main

import "fmt"

func reverseVowels(s string) string {
	n := len(s)
	arr := []byte(s)

	left, right := 0, n-1

	for left < right {
		for left < right {
			if isVowels(arr[left]) == false {
				left ++
			}else {
				break
			}
		}

		for left < right {
			if isVowels(arr[right]) == false {
				right--
			}else {
				break
			}
		}

		if left == right {
			break
		}

		arr[left], arr[right] = arr[right], arr[left]
		left ++
		right --
	}

	return string(arr)
}

func isVowels(a byte) bool {
	if a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' ||a == 'A' || a == 'E' || a == 'I' || a == 'O' || a == 'U'{
		return true
	}
	return false
}

func main() {
 	fmt.Println(reverseVowels("hello"))
}
