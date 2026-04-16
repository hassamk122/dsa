package main

/*
 Min Stack
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

    MinStack() initializes the stack object.
    void push(int val) pushes the element val onto the stack.
    void pop() removes the element on the top of the stack.
    int top() gets the top element of the stack.
    int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function

*/

type MinStack struct {
	mainArr []int
	minArr  []int
}

func Constructor() MinStack {
	return MinStack{
		mainArr: make([]int, 0),
		minArr:  make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if this.IsEmpty() {
		this.minArr = append(this.minArr, val)
	}

	if val <= this.GetMin() {
		this.minArr = append(this.minArr, val)
	}

	this.mainArr = append(this.mainArr, val)

}

func (this *MinStack) Pop() {
	if this.IsEmpty() {
		return
	}

	if this.Top() == this.GetMin() {
		this.minArr = this.minArr[:len(this.minArr)-1]
	}

	this.mainArr = this.mainArr[:len(this.mainArr)-1]
}

func (this *MinStack) IsEmpty() bool {
	return len(this.mainArr) == 0
}

func (this *MinStack) Top() int {
	if len(this.mainArr) == 0 {
		return -1
	}

	return this.mainArr[len(this.mainArr)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minArr) == 0 {
		return this.Top()
	}

	return this.minArr[len(this.minArr)-1]
}
