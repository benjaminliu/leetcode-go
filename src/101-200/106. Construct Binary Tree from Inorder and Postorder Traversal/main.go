package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	len1 := len(inorder)
	if len1 == 0 {
		return nil
	}

	maps := make(map[int]int)
	for i := 0; i < len1; i++ {
		maps[inorder[i]] = i
	}

	idx := len1 - 1

	return helper(postorder, &idx, maps, 0, len1-1)
}
func helper(postorder []int, idx *int, maps map[int]int, left int, right int) *TreeNode {
	root := &TreeNode{Val: postorder[*idx]}
	if left == right {
		return root
	}
	mid := maps[ postorder[*idx]]
	if mid < right {
		*idx --
		root.Right = helper(postorder, idx, maps, mid+1, right)
	}
	if left < mid {
		*idx --
		root.Left = helper(postorder, idx, maps, left, mid-1)
	}
	return root
}

func main() {
	root := buildTree([]int{1, 2, 3, 4}, []int{2, 1, 4, 3})
	if root == nil {
		fmt.Println("empty tree")
	} else {
		printTree(root)
	}
}

func printTree(root *TreeNode) {
	fmt.Printf("%v -> ", root.Val)
	if root.Left != nil {
		printTree(root.Left)
	}
	if root.Right != nil {
		printTree(root.Right)
	}
}
