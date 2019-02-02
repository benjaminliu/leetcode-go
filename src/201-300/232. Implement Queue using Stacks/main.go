package main

type MyQueue struct {
	input  []int
	output []int
	size   int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{input: make([]int, 0), output: make([]int, 0), size: 0}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
	this.size++
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			last := len(this.input) - 1
			this.output = append(this.output, this.input[last])
			this.input = this.input[0:last]
		}
	}
	last := len(this.output) - 1
	res := this.output[last]
	this.output = this.output[0:last]
	this.size--
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			last := len(this.input) - 1
			this.output = append(this.output, this.input[last])
			this.input = this.input[0:last]
		}
	}
	return this.output[ len(this.output)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {

}
