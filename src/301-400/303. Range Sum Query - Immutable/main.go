package main

import "fmt"

type NumArray struct {
	sums []int
	size int
}

func Constructor(nums []int) NumArray {
	len1 := len(nums)
	temp := make([]int, len1)
	if len1 == 0 {
		return NumArray{sums: temp, size: 0}
	}
	temp[0] = nums[0]
	for i := 1; i < len1; i++ {
		temp[i] = temp[i-1] + nums[i]
	}

	return NumArray{sums: temp, size: len(nums)}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 || j > this.size {
		return 0
	}
	if i == 0 {
		return this.sums[j]
	}

	return this.sums[j] - this.sums[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {

	//nums := []int{-1}

	nums := []int{-2, 0, 3, -5, 2, -1}
	//
	obj := Constructor(nums)

	fmt.Println(obj.SumRange(0, 0))
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))

}
