package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		temp := fast.Next
		fast.Next = temp.Next
		temp.Next = slow.Next
		slow.Next = temp
		fast = fast.Next
		slow = slow.Next
	}
	return head
}

func main() {

}
