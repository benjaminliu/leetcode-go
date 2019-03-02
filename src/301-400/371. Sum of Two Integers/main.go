package main

//bit,  need more care
func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {

}
