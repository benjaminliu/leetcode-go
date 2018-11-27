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
	cur := dummy
	for cur.Next != nil && cur.Next.Next!= nil{
		for cur.Next.Next!= nil && cur.Next.Val == cur.Next.Next.Val {
			cur.Next.Next = cur.Next.Next.Next
		}
		cur = cur.Next
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
