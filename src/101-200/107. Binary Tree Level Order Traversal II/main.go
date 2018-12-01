package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)

	helper(&res, root, 0)
	len1 := len(res)
	start := 0
	end := len1 - 1
	for start < end {
		res[start], res[end] = res[end], res[start]
		start++
		end --
	}
	return res
}
func helper(res *[][]int, node *TreeNode, depth int) {
	if depth == len(*res) {
		*res = append(*res, make([]int, 0))
	}
	(*res)[depth] = append((*res)[depth], node.Val)
	depth++
	if node.Left != nil {
		helper(res, node.Left, depth)
	}
	if node.Right != nil {
		helper(res, node.Right, depth)
	}
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	res := levelOrderBottom(root)

	fmt.Println(res)
}
