package main

func singleNumber(nums []int) int {

	maps := make(map[int]int)
	for _,value := range nums{
		if count ,ok := maps[value];ok{
			maps[value] = count+1
		}else {
			maps[value] = 1
		}
	}

	for key,value :=range maps{
		if value==1 {
			return key
		}
	}
	return -1
}

func singleNumber1(nums []int) int {
	a, b := 0, 0
	for _, v := range nums {
		b = (b ^ v) & (^a)
		a = (a ^ v) & (^b)
	}
	return b
}

func main() {
	
}
