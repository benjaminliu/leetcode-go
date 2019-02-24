package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return helper(root, make(map[*TreeNode]int))
}
func helper(root *TreeNode, maps map[*TreeNode]int) int {
	if root == nil {
		return 0
	}

	if value, ok := maps[root]; ok {
		return value
	}

	val := 0

	if root.Left != nil {
		val += helper(root.Left.Left, maps)
		val += helper(root.Left.Right, maps)
	}

	if root.Right != nil {
		val += helper(root.Right.Left, maps)
		val += helper(root.Right.Right, maps)
	}

	val = max(val+root.Val, helper(root.Left, maps)+helper(root.Right, maps))

	maps[root] = val
	return val
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}


func rob1(root *TreeNode) int {
	res := robSub(root);
	return max(res[0],res[1])
}
func robSub(root *TreeNode) []int {
	if root == nil {
		return make([]int,2)
	}

	left := robSub(root.Left)
	right := robSub(root.Right)

	res := make([]int,2)

	//do not rob root
	res[0] = max(left[0],left[1]) + max(right[0],right[1])

	//rob root
	res[1] = root.Val + left[0] +right[0]
	return res
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
}
