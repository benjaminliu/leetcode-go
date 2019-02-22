package main

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder,",")

	//at least 1 node
	diff :=1
	for _,value := range nodes{
		//this node count down
		diff --
		if diff<0 {
			return false
		}
		//a none # node has 2 children
		if value != "#"{
			diff+=2
		}
	}
	return diff ==0
}

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}
