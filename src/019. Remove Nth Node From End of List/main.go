package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	fast := dummy
	slow := dummy

	for i := n; i > 0; i-- {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			return dummy.Next
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	node := removeNthFromEnd(head, 2)
	for node != nil {
		fmt.Printf("%v ->", node.Val)

		node = node.Next
	}
}
