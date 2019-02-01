package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var res int
	idx := 0
	dfs(root, &idx, k, &res)
	return res
}
func dfs(root *TreeNode, idx *int, k int, res *int) bool {
	if root.Left != nil {
		if dfs(root.Left, idx, k, res) == true {
			return true
		}
	}
	*idx ++
	if *idx == k {
		*res = root.Val
		return true
	}
	if root.Right != nil {
		if dfs(root.Right, idx, k, res) == true {
			return true
		}
	}
	return false
}

func main() {
	//root := &TreeNode{Val: 3}
	//root.Left = &TreeNode{Val: 1}
	//root.Right = &TreeNode{Val: 4}
	//root.Left.Right = &TreeNode{Val: 2}

	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 1}

	fmt.Println(kthSmallest(root, 3))
}
