package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	preorder(&res, root, 0)
	return res
}
func preorder(res *[][]int, root *TreeNode, depth int) {
	if depth >= len(*res) {
		*res = append(*res, make([]int, 0))
	}

	(*res)[depth] = append((*res)[depth], root.Val)

	depth++
	if root.Left != nil {
		preorder(res, root.Left, depth)
	}
	if root.Right != nil {
		preorder(res, root.Right, depth)
	}
}

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	res = preorder1(res, root, 0)
	return res
}
func preorder1(res [][]int, root *TreeNode, depth int) [][]int {
	if depth >= len(res) {
		res = append(res, make([]int, 0))
	}

	res[depth] = append(res[depth], root.Val)

	depth++
	if root.Left != nil {
		res = preorder1(res, root.Left, depth)
	}
	if root.Right != nil {
		res = preorder1(res, root.Right, depth)
	}
	return res
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	res := levelOrder(root)

	fmt.Println(res)
}
