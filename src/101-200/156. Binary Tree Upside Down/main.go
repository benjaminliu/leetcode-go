package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	res := helper(root, root.Left, root.Right)
	root.Left = nil
	root.Right = nil
	return res
}

func helper(root *TreeNode, left *TreeNode, right *TreeNode) *TreeNode {
	leftLeft := left.Left
	leftRight := left.Right
	left.Right = root
	left.Left = right
	if leftLeft != nil {
		return helper(left, leftLeft, leftRight)
	}
	return left
}

func upsideDownBinaryTree1(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	newRoot := upsideDownBinaryTree(root.Left)
	right := root.Right
	newRoot.Right.Left = right
	newRoot.Right.Right = root
	root.Left = nil
	root.Right = nil
	return newRoot
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	res := upsideDownBinaryTree1(root)
	fmt.Println(res.Val)
}
