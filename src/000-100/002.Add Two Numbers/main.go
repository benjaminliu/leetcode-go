package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	cur := head

	for {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10
		cur.Val = sum % 10
		if l1 == nil && l2 == nil {
			break
		}
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return head
}

func main() {
	c := &ListNode{3, nil}
	b := &ListNode{2, c}
	a := &ListNode{1, b}
	y := &ListNode{8, nil}
	x := &ListNode{7, y}
	res := addTwoNumbers(a, x)
	for res != nil {
		fmt.Print(res.Val)
		fmt.Print("->")
		res = res.Next
	}
}
