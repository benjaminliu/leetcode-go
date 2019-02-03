package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	pPath := findPath(root, p, make([]*TreeNode, 0))
	qPath := findPath(root, q, make([]*TreeNode, 0))

	lenP := len(pPath)
	lenQ := len(qPath)

	if lenP == 1 {
		return p
	}
	if lenQ == 1 {
		return q
	}

	len1 := 0
	if lenP > lenQ {
		len1 = lenQ
	} else {
		len1 = lenP
	}

	for i := 0; i < len1; i++ {
		if pPath[i].Val != qPath[i].Val {
			return pPath[i-1]
		}
	}
	return pPath[len1-1]
}
func findPath(root *TreeNode, node *TreeNode, path []*TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	path = append(path, root)
	if root.Val == node.Val {
		return path
	}

	res := findPath(root.Left, node, path)
	if res != nil {
		return res
	}

	res = findPath(root.Right, node, path)
	return res
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val {
		return p
	}
	if root.Val == q.Val {
		return q
	}

	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	p := &TreeNode{Val: 5}
	q := &TreeNode{Val: 4}

	fmt.Println(lowestCommonAncestor1(root, p, q).Val)
}
