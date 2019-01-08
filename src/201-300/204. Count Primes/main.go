package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	noPrimes := make([]bool, n)
	noPrimes[0] = true
	noPrimes[1] = true
	sqrt := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt; i++ {
		if noPrimes[i] == false {
			for j := 2; j*i < n; j++ {
				noPrimes [j*i] = true
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if noPrimes[i] == false {
			count++
		}
	}
	return count
}

func countPrimes1(n int) int {
	n--
	if n < 2 {
		return 0
	}
	if n < 3 {
		return 1
	}
	if n < 5 {
		return 2
	}

	count := 2
	primes := []int{3}
outter:
	for i := 5; i <= n; i += 2 {
		for j := 0; j < len(primes); j++ {
			temp := primes[j]
			if i%temp == 0 {
				continue outter
			}
		}
		count++
		primes = append(primes, i)
	}
	return count
}

func main() {
	fmt.Println(countPrimes(10))
}
