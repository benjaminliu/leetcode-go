package main

func reverseString(s []byte) {
	n := len(s)
	if n < 2 {
		return
	}

	left, right := 0, n-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left ++
		right --
	}
}

func main() {

}
