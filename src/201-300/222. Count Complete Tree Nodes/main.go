package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	count := 0
	helper(root, &count)
	return count
}
func helper(root *TreeNode, count *int) {
	if root == nil {
		return
	}
	*count++
	helper(root.Left, count)
	helper(root.Right, count)
}

func main() {

}
