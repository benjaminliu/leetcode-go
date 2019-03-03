package main

import (
	"container/heap"
	"fmt"
)

type MinHeap [][]int

func (this MinHeap) Len() int {
	return len(this)
}

func (this MinHeap) Less(i, j int) bool {
	return (this[i][0] + this[i][1]) <= (this[j][0] + this[j][1])
}

func (this MinHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *MinHeap) Push(x interface{}) {
	(*this) = append((*this), x.([]int))
}

func (this *MinHeap) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 == 0 || n2 == 0 || k == 1 {
		return nil
	}

	maxHeap := &MinHeap{}
	for i := 0; i < n1 && i < k; i++ {
		maxHeap.Push([]int{nums1[i], nums2[0], 0})
	}
	heap.Init(maxHeap)

	res := make([][]int, 0)
	for k > 0 && maxHeap.Len() != 0 {
		k--

		cur := heap.Pop(maxHeap).([]int)
		res = append(res, []int{cur[0], cur[1]})
		if cur[2] == n2-1 {
			continue
		}
		heap.Push(maxHeap, []int{cur[0], nums2[cur[2]+1], cur[2] + 1})
	}
	return res
}

func main() {
	var num1 []int
	var num2 []int

	num1 = []int{1, 2, 4, 5, 6}
	num2 = []int{3, 5, 7, 9}

	fmt.Println(kSmallestPairs(num1, num2, 3))
	//
	//num1 = []int{1, 7, 11}
	//num2 = []int{2, 4, 6}
	//
	//fmt.Println(kSmallestPairs(num1, num2, 3))

	//num1 = []int{1, 1, 2}
	//num2 = []int{1, 2, 3}
	//
	//fmt.Println(kSmallestPairs(num1, num2, 2))
}
