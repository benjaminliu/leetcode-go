package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := math.MaxInt64
	helper(root, 0, &depth)
	return depth
}
func helper(node *TreeNode, parentDepth int, depth *int) {

	d := parentDepth + 1

	if node.Left == nil && node.Right == nil {
		if d < *depth {
			*depth = d
		}
	}

	if node.Left != nil {
		helper(node.Left, d, depth)
	}
	if node.Right != nil {
		helper(node.Right, d, depth)
	}
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if root.Left == nil {
		return right + 1
	}
	if root.Right == nil {
		return left + 1
	}

	if left > right {
		return right + 1
	} else {
		return left + 1
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	//root.Right = &TreeNode{Val:3}
	//root.Left.Left = &TreeNode{Val:4}
	//root.Right.Right = &TreeNode{Val:5}

	fmt.Println(minDepth(root))
}
