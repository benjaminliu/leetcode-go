package main

import "math/rand"

type Solution struct {
	arr [] int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.arr
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.arr)
	if n == 0 {
		return nil
	}
	temp := make([]int, n)
	copy(temp, this.arr)

	for i := 0; i < n; i++ {
		t := rand.Intn(i)
		temp[i], temp[t] = temp[t], temp[i]
	}
	return temp
}

func main() {

}
