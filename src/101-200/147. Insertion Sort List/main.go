package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: 0, Next: head}
	pre := head
	cur := head.Next

	for cur != nil {
		temp := dummy
		for temp.Next != nil && temp.Next != cur && temp.Next.Val < cur.Val {
			temp = temp.Next
		}
		if temp.Next == cur{
			cur = cur.Next
			pre = pre.Next
			continue
		}
		if temp.Next != nil {
			pre.Next = cur.Next
			cur.Next = temp.Next
			temp.Next = cur
			cur = pre.Next
		} else {
			cur = cur.Next
			pre = pre.Next
		}
	}

	return dummy.Next
}

func main() {
	//head := &ListNode{Val: 4}
	//head.Next = &ListNode{Val: 2}
	//head.Next.Next = &ListNode{Val: 1}
	//head.Next.Next.Next = &ListNode{Val: 3}

	head := &ListNode{Val: -1}
	head.Next = &ListNode{Val: 5}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val:0}

	res := insertionSortList(head)
	for res != nil {
		fmt.Printf("%v -> ", res.Val)
		res = res.Next
	}
}
