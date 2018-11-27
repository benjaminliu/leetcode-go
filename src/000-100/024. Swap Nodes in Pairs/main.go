package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	pre, cur, next := dummy, head, head.Next

	for {
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next

		if cur.Next == nil || cur.Next.Next == nil {
			break
		}

		pre = cur
		cur = cur.Next
		next = cur.Next
	}
	return dummy.Next
}

func swapPairs1(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre, cur := dummy, head
	for cur != nil && cur.Next != nil {
		pre.Next = cur.Next
		cur.Next = pre.Next.Next
		pre.Next.Next = cur
		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}

	res := swapPairs(head)
	for res != nil {
		fmt.Printf("%v-->", res.Val)
		res = res.Next
	}
}
