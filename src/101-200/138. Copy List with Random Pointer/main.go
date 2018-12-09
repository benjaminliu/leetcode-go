package main

import "fmt"

type RandomListNode struct {
	label  int
	next   *RandomListNode
	random *RandomListNode
}

func copyRandomList(node *RandomListNode) *RandomListNode {
	if node == nil {
		return nil
	}
	cur := node
	maps := make(map[int]*RandomListNode)
	last := &RandomListNode{label: cur.label}
	maps[last.label] = last
	cur = cur.next
	for cur != nil {
		temp := &RandomListNode{label: cur.label}
		maps[temp.label ] = temp
		last.next = temp
		last = last.next
		cur = cur.next
	}

	cur = node
	for cur != nil {
		if cur.random != nil {
			maps[cur.label].random = maps[cur.random.label]
		}
		cur = cur.next
	}
	return maps[node.label]
}

func main() {
	node1 := &RandomListNode{label: 1}
	node2 := &RandomListNode{label: 2}
	node3 := &RandomListNode{label: 3}
	node4 := &RandomListNode{label: 4}
	node1.next = node2
	node2.next = node3
	node3.next = node4

	node1.random = node4
	node3.random = node1
	node4.random = node4

	res := copyRandomList(node1)

	if res != nil {
		fmt.Printf("%v", res.label)
	}

}
