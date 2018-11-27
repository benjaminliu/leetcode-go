package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	dfs(root, &res)
	return res
}
func dfs(node *TreeNode, res *[]int) {
	if node.Left != nil {
		dfs(node.Left, res)
	}
	*res = append(*res, node.Val)
	if node.Right != nil {
		dfs(node.Right, res)
	}
}

func main() {

}
