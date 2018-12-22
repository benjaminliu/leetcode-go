package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lenA, lenB := getLen(headA), getLen(headB)

	for lenA > lenB {
		headA = headA.Next
		lenA--
	}
	for lenA < lenB {
		headB = headB.Next
		lenB--
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func getLen(node *ListNode) int {
	len1 := 0
	for node != nil {
		len1++
		node = node.Next
	}
	return len1
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func main() {

	headA := &ListNode{Val: 4}
	headA.Next = &ListNode{Val: 1}
	headA.Next.Next = &ListNode{Val: 8}
	headA.Next.Next.Next = &ListNode{Val: 4}
	headA.Next.Next.Next.Next = &ListNode{Val: 5}

	headB := &ListNode{Val: 5}
	headB.Next = &ListNode{Val: 0}
	headB.Next.Next = &ListNode{Val: 1}
	headB.Next.Next.Next = &ListNode{Val: 8}
	headB.Next.Next.Next.Next = &ListNode{Val: 4}
	headB.Next.Next.Next.Next.Next = &ListNode{Val: 5}

	res := getIntersectionNode1(headA, headB)

	if res != nil {
		fmt.Printf("%v", res.Val)
	}
}
