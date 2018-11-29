package _01__Symmetric_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSame(root.Left, root.Right)
}
func isSame(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	if isSame(left.Left, right.Right) == false {
		return false
	}
	if isSame(left.Right, right.Left) == false {
		return false
	}
	return true
}

func main() {

}
