package main

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	n    int
}

/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	temp := head
	n := 0
	for temp != nil {
		n++
		temp = temp.Next
	}
	return Solution{head: head, n: n}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	i := rand.Intn(this.n)
	temp := this.head
	for i > 0 {
		temp = temp.Next
		i--
	}
	return temp.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

func main() {

}
