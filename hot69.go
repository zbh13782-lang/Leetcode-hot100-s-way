package hot

import "math"

type MinStack struct {
	stack   []int
	mystack []int
}

func Constructor() MinStack {
	mystack := make([]int, 0)
	mystack = append(mystack, math.MaxInt)
	return MinStack{
		make([]int, 0),
		mystack,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if val <= this.mystack[len(this.mystack)-1] {
		this.mystack = append(this.mystack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	if this.Top() == this.mystack[len(this.mystack)-1] {
		this.stack = this.stack[:len(this.stack)-1]
		this.mystack = this.mystack[:len(this.mystack)-1]
	} else {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return math.MaxInt
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mystack[len(this.mystack)-1]
}