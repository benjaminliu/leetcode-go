package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

func connect(root *TreeNode) {
	if root == nil {
		return
	}

	parentQueue := list.New()
	childQueue := list.New()

	parentQueue.PushBack(root)

	for parentQueue.Len() != 0 {
		cur := parentQueue.Front()
		node := cur.Value.(*TreeNode)
		if node.Left != nil {
			childQueue.PushBack(node.Left)
		}
		if node.Right != nil {
			childQueue.PushBack(node.Right)
		}
		parentQueue.Remove(cur)
		if parentQueue.Len() == 0 {
			parentQueue, childQueue = childQueue, parentQueue
		} else {
			node.Next = parentQueue.Front().Value.(*TreeNode)
		}
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	connect(root)
	fmt.Printf("%v ->", root.Val)

}
