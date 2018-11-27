package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head

	var dup *ListNode
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			dup = cur.Next
		} else {
			dup = nil
			cur = cur.Next
		}
		if dup != nil {
			for cur.Next != nil && cur.Next.Val == dup.Next.Val {
				cur.Next = cur.Next.Next
			}
		}
	}
	return dummy.Next
}

func main() {
	//head := &ListNode{Val: 1}
	//head.Next = &ListNode{Val: 2}
	//head.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next.Next = &ListNode{Val: 4}
	//head.Next.Next.Next.Next.Next = &ListNode{Val: 5}

	//head := &ListNode{Val: 1}
	//head.Next = &ListNode{Val: 1}
	//head.Next.Next = &ListNode{Val: 1}
	//head.Next.Next.Next = &ListNode{Val: 2}
	//head.Next.Next.Next.Next = &ListNode{Val: 3}

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}


	//head := &ListNode{Val: 1}
	//head.Next = &ListNode{Val: 2}
	//head.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next.Next = &ListNode{Val: 4}
	//head.Next.Next.Next.Next.Next=&ListNode{Val:4}
	//head.Next.Next.Next.Next.Next.Next=&ListNode{Val:5}


	res := deleteDuplicates(head)

	for res != nil {
		fmt.Printf("%v->", res.Val)
		res = res.Next
	}
}
