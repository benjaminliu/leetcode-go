package main

//math problem
func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	if x+y < z {
		return false
	}

	if x == z || y == z || x+y == z {
		return true
	}

	return z%gcd(x, y) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func main() {
	//2, 7 , 4
}
