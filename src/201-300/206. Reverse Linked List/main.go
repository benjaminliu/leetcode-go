package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}

	var last,cur,next *ListNode
	cur = head

	for cur != nil{
		next = cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	return last
}

func main() {

}
