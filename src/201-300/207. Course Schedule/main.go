package main

import (
	"container/list"
	"fmt"
)

//just ignore
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	inDegree := make([]int, numCourses)
	len1 := len(prerequisites)

	for i := 0; i < len1; i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
		inDegree[prerequisites[i][0]]++
	}
	queue := list.New()
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}

	count := 0
	for queue.Len() != 0 {
		count++
		temp := queue.Front()
		queue.Remove(temp)
		cur := temp.Value.(int)
		len2 := len(graph[cur])
		for i := 0; i < len2; i++ {
			next := graph[cur][i]
			inDegree[next]--
			if inDegree[next] == 0 {
				queue.PushBack(next)
			}
		}
	}

	return count == numCourses
}

func main() {
	prerequisites := [][]int{{1, 0}}
	fmt.Println(canFinish(2, prerequisites))
}
