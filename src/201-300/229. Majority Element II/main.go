package main

import "fmt"

func majorityElement(nums []int) []int {
	len1 := len(nums)
	if len1 == 0 {
		return nil
	}
	res := make([]int, 0)
	if len1 == 1 {
		res = append(res, nums[0])
		return res
	}
	if len1 == 2 {
		if nums[0] == nums[1] {
			res = append(res, nums[0])
			return res
		}
		res = append(res, nums[0])
		res = append(res, nums[1])
		return res
	}

	maps := make(map[int]int)

	for _, value := range nums {
		if value1, ok := maps[value]; ok {
			maps[value ] = value1 + 1
		} else {
			maps[value] = 1
		}
	}
	target := float64(len1) / 3

	for k, v := range maps {
		if float64(v) > target {
			res = append(res, k)
		}
	}
	return res
}

func majorityElement1(nums []int) []int {
	len1 := len(nums)
	if len1 == 0 {
		return nil
	}
	res := make([]int, 0)
	if len1 == 1 {
		res = append(res, nums[0])
		return res
	}
	if len1 == 2 {
		if nums[0] == nums[1] {
			res = append(res, nums[0])
			return res
		}
		res = append(res, nums[0])
		res = append(res, nums[1])
		return res
	}

	num1, num2 := 0, 0
	cnt1, cnt2 := 0, 0
	for _,v := range nums{
		if v == num1{
			cnt1++
		}else if v == num2{
			cnt2++
		}else if cnt1 ==0{
			num1 = v
			cnt1 = 1
		}else  if cnt2 ==0{
			num2 = v
			cnt2 = 1
		}else {
			cnt1--
			cnt2--
		}
	}

	cnt1,cnt2 = 0,0
	for _,v := range nums{
		if v == num1{
			cnt1++
		}else if(v == num2){
			cnt2++
		}
	}
	target := float64(len1) / 3
	if float64(cnt1) > target{
		res = append(res, num1)
	}
	if float64(cnt2) > target{
		res = append(res, num2)
	}
	return res
}

func main() {
	fmt.Println(majorityElement1([]int{1}))
	fmt.Println(majorityElement1([]int{1, 2}))
	fmt.Println(majorityElement1([]int{3, 2, 3}))
	fmt.Println(majorityElement1([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
