package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		if isSameTree(p.Left, q.Left) == false {
			return false
		}
		if isSameTree(p.Right, q.Right) == false {
			return false
		}
		return true
	}
	return false
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}

	root1 := &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 2}

	fmt.Println(isSameTree(root, root1))
}
