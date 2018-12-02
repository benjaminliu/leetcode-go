package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	group := make([]int, 0)
	for head != nil {
		group = append(group, head.Val)
		head = head.Next
	}

	return helper(group, 0, len(group)-1)
}
func helper(group []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{Val: group[mid]}
	root.Left = helper(group, left, mid-1)
	root.Right = helper(group, mid+1, right)
	return root
}

func main() {

}
