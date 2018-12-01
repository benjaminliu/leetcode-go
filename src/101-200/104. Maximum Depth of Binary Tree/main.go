package main

  type TreeNode struct {
	      Val int
	      Left *TreeNode
	      Right *TreeNode
	  }

func maxDepth(root *TreeNode) int {
	return dfs(root,0);
}
func dfs(node *TreeNode, depth int) int {
	if node== nil{
		return depth
	}
	depth++
	left := dfs(node.Left,depth)
	right := dfs(node.Right,depth)
	if left >=right{
		return left
	}
	return right
}

func main() {
	
}
