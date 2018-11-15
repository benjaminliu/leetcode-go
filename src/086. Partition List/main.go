package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head

	preBigger := dummy
	firstBigger := dummy.Next

	for firstBigger != nil {
		if firstBigger.Val >= x {
			break
		}
		firstBigger = firstBigger.Next
		preBigger = preBigger.Next
	}
	if firstBigger == nil {
		return head
	}

	cur := firstBigger.Next
	preCur := firstBigger

	var temp *ListNode
	for cur != nil {
		if cur.Val < x {
			temp = cur
			preCur.Next = cur.Next
			cur = cur.Next
			preBigger.Next = temp
			temp.Next = firstBigger
			preBigger= temp
		} else {
			cur = cur.Next
			preCur = preCur.Next
		}
	}
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1}
	//head.Next = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}

	res := partition(head, 3)
	for res != nil {
		fmt.Printf("%v -> ", res.Val)
		res = res.Next
	}
}
