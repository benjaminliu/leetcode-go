package main

import (
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses <= 0 {
		return nil
	}

	if prerequisites == nil || len(prerequisites) == 0 {

		res := make([]int, 0)
		for i := 0; i < numCourses; i++ {
			res = append(res, i)
		}
		return res
	}

	graph := getGraph(prerequisites, numCourses)

	visited := make([]int, numCourses)
	order := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if hasCycle(&order, &graph, &visited, i) {
			return nil
		}
	}
	res := make([]int, len(order))
	idx := 0
	for _, value := range order {
		res[idx] = value
		idx++
	}
	return res
}
func hasCycle(order *[]int, graph *[][]int, visited *[]int, i int) bool {
	if (*visited)[i] == 2 {
		return false
	}
	if (*visited)[i] == 1 {
		return true
	}
	(*visited)[i] = 1
	list := (*graph)[i]
	for _, value := range list {
		if hasCycle(order, graph, visited, value) {
			return true
		}
	}
	(*visited)[i] = 2
	*order = append(*order, i)
	return false
}
func getGraph(prerequistites [][]int, numCourses int) [][]int {
	res := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		res[i] = make([]int, 0)
	}
	len1 := len(prerequistites)
	for i := 0; i < len1; i++ {
		n0 := prerequistites[i][0]
		n1 := prerequistites[i][1]
		res[n0] = append(res[n0], n1)
	}
	return res
}

//this cannot work,
func findOrder1(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return nil
	}

	res := make([]int, 0)

	if prerequisites == nil || len(prerequisites) == 0 {
		for i := 0; i < numCourses; i++ {
			res = append(res, i)
		}
		return res
	}

	followMap := make(map[int]int)
	preMap := make(map[int][]int)
	queue := make(map[int]bool)

	for _, pair := range prerequisites {
		followMap[pair[1]] = pair[0]
		if list, ok := preMap[pair[0]]; ok {
			list = append(list, pair[1])
			preMap[pair[0]] = list
		} else {
			list := []int{pair[1]}
			preMap[pair[0]] = list
		}
		queue[pair[0]] = true
	}

	addedMap := make(map[int]bool)
	for len(queue) != 0 {
		var temp int
		for temp, _ = range queue {
			if _, ok := followMap[temp]; ok {
				continue
			}

			res = append(res, temp)
			addedMap[temp] = true
			list, _ := preMap[temp]
			for _, value := range list {
				delete(followMap, value)
			}
			break
		}
		delete(queue, temp)
	}
	//this is nonsense
	if len(res) != len(preMap) {
		return nil
	}

	for i := 0; i < numCourses; i++ {
		if _, ok := addedMap[i]; ok == false {
			res = append(res, i)
		}
	}
	reverse(res)
	return res
}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	//groups := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2},}
	//groups := [][]int{{1, 0} ,}
	groups := [][]int{{0, 2}, {1, 2}, {2, 0},}

	fmt.Println(findOrder(4, groups))
}
