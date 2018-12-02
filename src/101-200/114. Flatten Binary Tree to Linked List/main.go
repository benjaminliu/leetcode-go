package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	for cur.Left != nil || cur.Right != nil {
		if cur.Left != nil {
			temp := cur.Right
			cur.Right = cur.Left
			cur.Left = nil
			temp1 := cur.Right
			for temp1.Right != nil {
				temp1 = temp1.Right
			}
			temp1.Right = temp
		}
		cur = cur.Right
	}
}

func flatten1(root *TreeNode) {
	for ; root != nil; root = root.Right {
		right := root.Right
		root.Left, root.Right = nil, root.Left

		p := root
		for p.Right != nil {
			p = p.Right
		}
		p.Right = right
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}

	flatten(root)
	for root != nil {
		fmt.Printf("%v -> ", root.Val)
		root = root.Right
	}
}
