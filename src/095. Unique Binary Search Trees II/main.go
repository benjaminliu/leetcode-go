package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n <=0 {
		return nil
	}

	return dfs(1,n)
}
func dfs(start int, end int) []*TreeNode {
	res := make([]*TreeNode,0)
	if start > end {
		return append(res, nil)
	}

	for i := start ;i <=end;i++{
		leftNode := dfs(start,i-1)
		rightNode := dfs(i+1, end)
		for _,left := range leftNode{
			for  _,right := range rightNode{
				root := new (TreeNode)
				root.Val= i
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	return res
}

func main() {
	 res := generateTrees(3)
	 for _,value := range res{
	 	fmt.Printf("%v -> " ,value.Val)
	 }


}
