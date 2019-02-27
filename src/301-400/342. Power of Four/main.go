package main

func isPowerOfFour(num int) bool {
	if num == 1 {
		return true
	}
	if num < 4 {
		return false
	}
	base := 4
	for num%4 == 0 {
		base = 4
		for num%base == 0 {
			num = num / base
			base = base * base
		}
	}
	return num ==1
}

func main() {

}
