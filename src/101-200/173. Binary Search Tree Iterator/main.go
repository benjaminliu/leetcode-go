package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	idx int
	arr []int
}

func Constructor(root *TreeNode) BSTIterator {
	arr := make([]int, 0)
	if root != nil {
		dfs(root, &arr)
	}
	return BSTIterator{idx: 0, arr: arr}
}
func dfs(node *TreeNode, arr *[]int) {
	if node.Left != nil {
		dfs(node.Left, arr)
	}

	*arr = append(*arr, node.Val)

	if node.Right != nil {
		dfs(node.Right, arr)
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	res := this.arr[this.idx ]
	this.idx++
	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.idx >= len(this.arr) {
		return false
	}
	return true
}

func main() {

}
