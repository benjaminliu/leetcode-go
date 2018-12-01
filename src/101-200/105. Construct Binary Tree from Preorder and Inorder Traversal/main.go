package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	len1 := len(preorder)
	if len1 == 0 {
		return nil
	}

	maps := make(map[int]int)
	for i := 0; i < len1; i++ {
		maps[inorder[i]] = i
	}

	idx :=0
	return helper(preorder, &idx, maps, 0, len1-1)
}
func helper(preorder []int, idx *int, maps map[int]int, left int, right int) *TreeNode {
	root := &TreeNode{Val: preorder[*idx]}
	if left == right {
		return root
	}
	mid := maps[ preorder[*idx]]
	if left < mid {
		*idx ++
		root.Left = helper(preorder, idx, maps, left, mid-1)
	}
	if mid < right {
		*idx ++
		root.Right = helper(preorder, idx, maps, mid+1, right)
	}
	return root
}

func main() {
	root := buildTree([]int{7,-10,-4,3,-1,2,-8,11}, []int{-4,-10,3,-1,7,11,-8,2})
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
