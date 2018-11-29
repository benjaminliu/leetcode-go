package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)

	dfs(&res, root, 0, )
	return res
}
func dfs(res *[][]int, root *TreeNode, depth int) {
	if depth >= len(*res) {
		*res = append(*res, make([]int, 0))
	}

	if depth%2 == 0 {
		(*res)[depth] = append((*res)[depth], root.Val)
	} else {
		(*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
	}

	depth++
	if root.Left != nil {
		dfs(res, root.Left, depth)
	}
	if root.Right != nil {
		dfs(res, root.Right, depth)
	}
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	res := zigzagLevelOrder(root)

	fmt.Println(res)
}
