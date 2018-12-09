package main

import "fmt"

type UndirectedGraphNode struct {
	label     int
	neighbors []*UndirectedGraphNode
}

func cloneGraph(node *UndirectedGraphNode) *UndirectedGraphNode {
	if node == nil {
		return nil
	}

	maps := make(map[int]*UndirectedGraphNode)

	return dfs(node, maps)
}
func dfs(node *UndirectedGraphNode, maps map[int]*UndirectedGraphNode) *UndirectedGraphNode {
	var dup *UndirectedGraphNode
	var ok bool
	if dup, ok = maps[node.label]; ok {
		//already exist， just return
	} else {
		dup = &UndirectedGraphNode{label: node.label}
		//need to set maps first, or will cause infinite loop
		maps[dup.label ] = dup

		len1 := len(node.neighbors)
		dup.neighbors = make([]*UndirectedGraphNode, len1)
		for i := 0; i < len1; i++ {
			dup.neighbors[i] = dfs(node.neighbors[i], maps)
		}
	}

	return dup
}

func main() {
	node0 := &UndirectedGraphNode{label: 0}
	node1 := &UndirectedGraphNode{label: 1}
	node2 := &UndirectedGraphNode{label: 2}
	node0.neighbors = []*UndirectedGraphNode{node1, node2}
	node1.neighbors = []*UndirectedGraphNode{node2}
	node2.neighbors = []*UndirectedGraphNode{node2}

	dup := cloneGraph(node0)
	if dup != nil {
		fmt.Printf("%v ", dup.label)
	}
}
