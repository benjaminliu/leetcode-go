package main

import "container/list"

type MyStack struct {
	emptyA bool
	queueA *list.List
	queueB *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{emptyA: true, queueA: list.New(), queueB: list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.emptyA {
		this.queueA.PushBack(x)
		for this.queueB.Len() > 0 {
			item := this.queueB.Front()
			this.queueB.Remove(item)
			this.queueA.PushBack(item.Value)
		}
	} else {
		this.queueB.PushBack(x)
		for this.queueA.Len() > 0 {
			item := this.queueA.Front()
			this.queueA.Remove(item)
			this.queueB.PushBack(item.Value)
		}
	}
	this.emptyA = !this.emptyA
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.emptyA {
		item := this.queueB.Front()
		this.queueB.Remove(item)
		return item.Value.(int)
	}
	item := this.queueA.Front()
	this.queueA.Remove(item)
	return item.Value.(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.emptyA {
		item := this.queueB.Front()
		return item.Value.(int)
	}
	item := this.queueA.Front()
	return item.Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.emptyA {
		return this.queueB.Len() == 0
	}
	return this.queueA.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {

}
