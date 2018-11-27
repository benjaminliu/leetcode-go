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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs(root, math.MaxInt64, math.MinInt64)
}

func dfs(node *TreeNode, max int, min int) bool {
	if node.Val <= min || node.Val >= max {
		return false
	}
	if node.Left != nil {
		if (node.Val > node.Left.Val && dfs(node.Left, node.Val, min)) == false {
			return false
		}
	}
	if node.Right != nil {
		if (node.Val < node.Right.Val && dfs(node.Right, max, node.Val)) == false {
			return false
		}
	}
	return true
}

func isValidBST1(root *TreeNode) bool {
	return isValidBst1(root, nil, nil)
}

func isValidBst1(root *TreeNode, l *TreeNode, r *TreeNode) bool {
	if root == nil {
		return true
	}

	if l != nil && root.Val <= l.Val {
		return false
	}

	if r != nil && root.Val >= r.Val {
		return false
	}

	return isValidBst1(root.Left, l, root) && isValidBst1(root.Right, root, r)
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	//root.Left.Left = &TreeNode{Val: 0}
	//root.Left.Right = &TreeNode{Val: 2}
	//root.Right.Left = &TreeNode{Val: 3}
	//root.Right.Right = &TreeNode{Val: 6}
	fmt.Println(isValidBST(root))
}
