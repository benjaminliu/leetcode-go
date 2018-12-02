package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return helper(root, 0, sum)
}
func helper(node *TreeNode, sum int, target int) bool {
	sum = sum + node.Val
	if node.Left == nil && node.Right == nil {
		if sum == target {
			return true
		}
		return false
	}
	if node.Left != nil {
		if helper(node.Left, sum, target) {
			return true
		}
	}
	if node.Right != nil {
		if helper(node.Right, sum, target) {
			return true
		}
	}
	return false
}

func main() {

}
