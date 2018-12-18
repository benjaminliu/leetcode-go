package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next ==nil{
		return head
	}

	mid := getMid(head)
	next := mid.Next
	mid.Next = nil

	return merge(sortList(head), sortList(next))
}

func merge(list *ListNode, list2 *ListNode) *ListNode {
	if list == nil {
		return list2
	}
	if list2 == nil {
		return list
	}

	var head *ListNode
	if list.Val <= list2.Val {
		head = list
		list = list.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	cur := head
	for list != nil && list2 != nil {
		if list.Val <=list2.Val {
			cur.Next = list
			list = list.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list != nil {
		cur.Next = list
	}
	if list2 != nil {
		cur.Next = list2
	}
	return head
}

func getMid(node *ListNode) *ListNode {
	if node == nil{
		return nil
	}
	slow, fast := node, node
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}
	return slow
}

func main() {
	head := &ListNode{Val: -1}
	head.Next = &ListNode{Val: 5}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 0}

	res := sortList(head)
	for res != nil {
		fmt.Printf("%v - > ", res.Val)
		res = res.Next
	}
}
