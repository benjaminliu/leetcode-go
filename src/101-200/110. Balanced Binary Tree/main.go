package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, ok := helper(root)
	return ok
}
func helper(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, ok1 := helper(node.Left)
	if ok1 == false {
		return -1, false
	}
	rightHeight, ok2 := helper(node.Right)
	if ok2 == false {
		return -1, false
	}

	if leftHeight >= rightHeight {
		diff := leftHeight - rightHeight
		if diff < 2 {
			return leftHeight + 1, true
		}
	} else {
		diff := rightHeight - leftHeight
		if diff < 2 {
			return rightHeight + 1, true
		}
	}

	return -1, false
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Right .Right = &TreeNode{Val:1}

	fmt.Println(isBalanced(root))
}
