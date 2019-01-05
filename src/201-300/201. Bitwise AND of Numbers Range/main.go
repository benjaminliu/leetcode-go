package main

func rangeBitwiseAnd(m int, n int) int {
	var i uint32
	for m != n {
		m >>= 1
		n >>= 1
		i++
	}
	return   m << i
}

func main() {

}
