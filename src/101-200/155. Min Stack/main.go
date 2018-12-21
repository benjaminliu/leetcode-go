package main

import "fmt"

type MinStack struct {
	val  int
	next *MinStack
	min  *MinStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{next: nil, min: nil}
}

func (this *MinStack) Push(x int) {
	oldTop := this.next
	this.next = &MinStack{val: x, next: oldTop}
	oldMin := this.min
	if oldMin == nil {
		this.min = &MinStack{val: x}
	} else {
		if x <= oldMin.val {
			this.min = &MinStack{val: x, next: oldMin}
		}
	}
}

func (this *MinStack) Pop() {
	oldTop := this.next
	if oldTop!= nil{
		this.next = oldTop.next

		if oldTop.val == this.min.val{
			this.min = this.min.next
		}
	}
}

func (this *MinStack) Top() int {
	if this.next != nil{
		return this.next.val
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.min != nil {
		return this.min.val
	}
	return 0
}

func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(4)
	obj.Push(2)
	obj.Push(2)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())

}
