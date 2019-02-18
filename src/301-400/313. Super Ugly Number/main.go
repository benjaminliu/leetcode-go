package main

import "math"

//need more care
func nthSuperUglyNumber(n int, primes []int) int {
	len1 := len(primes)
	idx := make([]int, len1)
	val := make([]int, len1)
	ugly := make([]int, n)

	for i := 0; i < len1; i++ {
		val[i] = 1
	}

	next := 1
	for i := 0; i < n; i++ {
		ugly [i] = next

		next = math.MaxInt32

		for j := 0; j < len1; j++ {
			if val[j] == ugly[i] {
				val[j] = ugly[idx[j]] * primes[j]
				idx[j] ++
			}
			next = minInt(next, val[j])
		}
	}
	return ugly[n-1]
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {

}
