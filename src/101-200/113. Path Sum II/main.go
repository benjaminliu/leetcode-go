package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	group := make([]int, 0)
	helper(root, group, &res, 0, sum)
	return res
}
func helper(node *TreeNode, group []int, res *[][]int, sum int, target int) {
	sum += node.Val
	group = append(group, node.Val)
	if node.Left == nil && node.Right == nil {
		if sum == target {
			temp :=make([]int,len(group))
			copy(temp,group)
			*res = append(*res, temp)
		}
		return
	}
	if node.Left != nil {
		helper(node.Left, group, res, sum, target)
	}
	if node.Right != nil {
		helper(node.Right, group, res, sum, target)
	}
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 1}

	fmt.Println(pathSum(root, 22))
}
