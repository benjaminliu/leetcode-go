package main

func missingNumber(nums []int) int {
	len1 := len(nums)
	maps := make([]int,len1 +1 )
	for _,value := range nums{
		maps[value] = 1
	}

	for idx,value:=range maps{
		if value ==0 {
			return idx
		}
	}

	return len1
}

func missingNumber1(nums []int) int {
	res := 0

	for idx,value :=range nums{
		res = res ^ idx^value
	}
	return res ^ len(nums)
}

func main() {
	
}
