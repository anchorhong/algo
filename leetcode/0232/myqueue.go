package _232

type MyQueue struct {
	inStack  []int
	outStack []int
	size     int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	var inStack []int
	var outStack []int
	return MyQueue{inStack, outStack, 0}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
	this.size++
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.size == 0 {
		panic("queue is empty")
	}
	if len(this.outStack) > 0 {
		x := this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[0 : len(this.outStack)-1]
		this.size--
		return x
	} else {
		for i := len(this.inStack) - 1; i > 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
		}
		this.size--
		x := this.inStack[0]
		this.inStack = []int{}
		return x
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.size == 0 {
		panic("queue is empty")
	}
	if len(this.outStack) > 0 {
		return this.outStack[len(this.outStack)-1]
	} else {
		return this.inStack[0]
	}
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
