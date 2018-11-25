package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}

	idx := 1
	last := dummy
	cur := dummy.Next
	for idx != m {
		cur = cur.Next
		last = last.Next
		idx ++
	}
	end := cur
	next := cur.Next

	for idx != n {
		temp := cur
		cur = next
		next = cur.Next
		cur.Next = temp
		idx++
	}
	end.Next = next
	last.Next = cur
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	res := reverseBetween(head, 2, 4)
	for res != nil {
		fmt.Printf("%v -> ", res.Val)
		res = res.Next
	}
}
