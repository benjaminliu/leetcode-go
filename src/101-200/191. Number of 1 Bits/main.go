package main

func hammingWeight(num uint32) int {
	res :=0
	for num > 0{
		temp := num%2
		if temp ==1 {
			res++
		}
		num >>=1
	}
	return res
}

func main() {
	
}
