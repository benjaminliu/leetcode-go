package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0)

	group := make([]byte, n*2)
	helper(&res, &group, 0, 0, 0, n)
	return res
}

func helper(res *[]string, group *[]byte, idx int, openCount int, closeCount int, n int) {
	if idx == n*2 {
		*res = append(*res, string(*group))
		return
	}
	if openCount < n {
		(*group)[idx] = '('
		helper(res, group, idx+1, openCount+1, closeCount, n)
	}
	//close count cannot bigger than open count
	if closeCount < openCount {
		(*group)[idx] = ')'
		helper(res, group, idx+1, openCount, closeCount+1, n)
	}
}

func generateParenthesis1(n int) []string {
	res := make([]string, 0)

	helper1(&res, "", 0, 0, n)
	return res
}

func helper1(res *[]string, cur string, openCount int, closeCount int, n int) {
	if len(cur) == n*2 {
		*res = append(*res, cur)
		return
	}
	if openCount < n {
		helper1(res, cur+"(", openCount+1, closeCount, n)
	}
	//close count cannot bigger than open count
	if closeCount < openCount {
		helper1(res, cur+")", openCount, closeCount+1, n)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
