package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	dummy := &ListNode{Val:0,Next:head}
	slow, fast := dummy, dummy

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}

	last := slow.Next
	slow.Next = nil

	cur := last.Next
	last.Next = nil
	var next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = last
		last = cur
		cur = next
	}

	fast = head
	slow = last

	cur = dummy
	for fast!= nil && slow!= nil{
		cur.Next = fast
		fast = fast.Next
		cur = cur.Next
		cur.Next = slow
		slow = slow.Next
		cur = cur.Next
	}
	if fast!= nil{
		cur.Next = fast
	}

	fmt.Println("%v", head.Val)
}

func main() {
	head :=&ListNode{Val:1}
	head.Next = &ListNode{Val:2}
	head.Next.Next = &ListNode{Val:3}
	head.Next.Next.Next = &ListNode{Val:4}
	head.Next.Next.Next.Next = &ListNode{Val:5}

	reorderList(head)
}
