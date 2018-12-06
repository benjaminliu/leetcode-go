package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0

	helper(root, 0, &sum)
	return sum
}
func helper(node *TreeNode, temp int, sum *int) {
	if node == nil {
		return
	}
	temp = temp*10 + node.Val
	if node.Left == nil && node.Right == nil{
		*sum += temp
		return
	}

	if node.Left!= nil{
		helper(node.Left,temp,sum)
	}
	if node.Right!= nil{
		helper(node.Right,temp,sum)
	}
}

func main() {

}
