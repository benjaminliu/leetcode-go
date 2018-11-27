package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	if head == nil {
		return nil
	}
	length := 0
	dummy := new(ListNode)
	dummy.Next = head

	slow := dummy
	fast := dummy

	for k > 0 && fast.Next != nil {
		fast = fast.Next
		k--
		length++
	}
	if fast.Next == nil {
		if k == 0 {
			return head
		}
		k = k % length

		if k == 0 {
			return head
		}

		fast = dummy
		for k > 0 {
			fast = fast.Next
			k--
		}
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	temp := dummy.Next
	dummy.Next = slow.Next
	fast.Next = temp
	slow.Next = nil
	return dummy.Next
}

func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	slow := &ListNode{
		Next: head,
	}
	fast := &ListNode{
		Next: head,
	}
	steps := 0
	for steps < k {
		fast = fast.Next
		steps++
		if fast.Next == nil {
			k = k % steps
			if k == 0 {
				return head
			}
			steps = 0
			fast = slow

		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	newhead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newhead
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next = &ListNode{Val: 4}
	//head.Next.Next.Next.Next = &ListNode{Val: 5}

	res := rotateRight(head, 4)
	for res != nil {
		fmt.Printf("%v -> ", res.Val)
		res = res.Next
	}
}
