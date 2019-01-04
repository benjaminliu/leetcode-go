package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil{
		return nil
	}
	group := make([]int,0)
	parent := list.New()
	child := list.New()

	parent.PushBack(root)
	for parent.Len()>0{
		front := parent.Front()
		parent.Remove(front)
		temp:= front.Value.(*TreeNode)
		if temp.Left!= nil{
			child.PushBack(temp.Left)
		}
		if temp.Right!= nil{
			child.PushBack(temp.Right)
		}
		if(parent.Len() == 0 ){
			group = append(group, temp.Val)
			parent,child = child,parent
		}
	}
	return group
}

func main() {

}
