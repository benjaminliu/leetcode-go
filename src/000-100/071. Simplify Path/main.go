package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	len1 := len(path)
	if len1 < 2 {
		return "/"
	}

	strs := make([]string, 0)
	last := 0
	for i := 1; i < len1; i++ {
		if path[i] == '/' {
			if last+1 != i {
				strs = append(strs, string(path[last+1:i]))
			}
			last = i
		}
	}
	if last+1 != len1 {
		strs = append(strs, string(path[last+1:]))
	}
	if len(strs) == 0 {
		return "/"
	}

	res := make([]string, 0)
	for _, s := range strs {
		if s == "." {

		} else if s == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, s)
		}
	}
	if len(res) == 0 {
		return "/"
	}

	builder := strings.Builder{}
	for _, s := range res {
		builder.WriteString("/")
		builder.WriteString(s)
	}
	return builder.String()
}

func simplifyPath1(path string) string {
	split := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, x := range split {
		switch x {
		case "":
		case ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, x)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/../"))
}
