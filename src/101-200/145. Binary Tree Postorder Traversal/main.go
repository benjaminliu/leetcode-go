package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	group := make([]int, 0)
	helper(root, &group)
	return group
}
func helper(node *TreeNode, group *[]int) {
	if node.Left != nil {
		helper(node.Left, group)
	}
	if node.Right != nil {
		helper(node.Right, group)
	}
	*group = append(*group, node.Val)
}

func main() {

}
