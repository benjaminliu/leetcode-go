package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := make([]string, 0)
	helper(&res, root, make([]byte, 0))
	return res
}
func helper(res *[]string, node *TreeNode, group []byte) {
	group = append(group, strconv.Itoa(node.Val)...)

	if node.Left != nil || node.Right != nil {
		group = append(group, '-')
		group = append(group, '>')
	} else {
		*res = append(*res, string(group))
		return
	}
	if node.Left != nil {
		helper(res, node.Left, group)
	}
	if node.Right != nil {
		helper(res, node.Right, group)
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(binaryTreePaths(root))
}
