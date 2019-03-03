package main

//need more care

//  (a*b)%k = ((a%k)*(b%k))%k
func superPow(a int, b []int) int {
	n := len(b)
	if n == 0 {
		return 1
	}

	p := 0
	for _, value := range b {
		p = (p*10 + value) % 1140
	}

	if p == 0 {
		p += 1440
	}

	return power(a, p, 1337)
}
func power(a int, n int, mod int) int {
	a %= mod
	ret := 1
	for n != 0 {
		if (n & 1) != 0 {
			ret = ret * a % mod
		}
		a = a * a % mod
		n >>= 1
	}
	return ret
}

var base int = 1337

//  (a*b)%k = ((a%k)*(b%k))%k
// a^1234567 % k = (a^1234560 % k) * (a^7 % k) % k
// a^1234560 % k = ((a^123456 % k)^10)% k
func superPow1(a int, b []int) int {
	n := len(b)
	if n == 0 {
		return 1
	}

	lastDigit := b[n-1]

	return powMod(superPow(a, b[:n-1]), 10) * powMod(a, lastDigit) % base
}

//a^k mod 1337 where 0 <= k <= 10
func powMod1(a int, b int) int {
	a %= base
	res := 1
	for i := 0; i < b; i++ {
		res = (res * a) % base
	}
	return res
}

//a^k mod 1337 where 0 <= k <= 10
func powMod(a int, b int) int {
	if b == 0 {
		return 1
	}
	res := 1
	for i := 0; i < b; i++ {
		res = (res * (a % base)) % base
	}
	return res
}

//suppose:
//a = Ak +B;
//b = Ck + D;
//
//so:
//a * b = ACk + ADk + BCk +BD;
//
//ab % k = (ACk + ADk + BCk +BD) % k = BD % k;
//
//a % k = B;
//
//b % k = D;
//
//(a%k)(b%k)%k = BD % k;
//
//so we have:
//
//ab % k = (a%k)(b%k)%k

func main() {

}
