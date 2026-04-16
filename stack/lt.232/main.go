package main

/*
Implement queue using stacks

Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

	void push(int x) Pushes element x to the back of the queue.
	int pop() Removes the element from the front of the queue and returns it.
	int peek() Returns the element at the front of the queue.
	boolean empty() Returns true if the queue is empty, false otherwise.
*/
type Stack struct {
	arr []int
}

func (s *Stack) GetTop() int {
	if s.IsEmpty() {
		return int(-1)
	}
	return s.arr[len(s.arr)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) Push(data int) {
	s.arr = append(s.arr, data)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return int(-1)
	}
	lastValue := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return lastValue
}

type Queue struct {
	leftStack  *Stack
	rightStack *Stack
}

func NewQueue() *Queue {
	return &Queue{
		leftStack:  &Stack{},
		rightStack: &Stack{},
	}
}

func (q *Queue) IsEmpty() bool {
	return q.leftStack.IsEmpty() && q.rightStack.IsEmpty()
}

func (q *Queue) Enqueue(val int) {
	q.leftStack.Push(val)
}

func (q *Queue) Dequeue() int {
	if !q.rightStack.IsEmpty() {
		return q.rightStack.Pop()
	}

	for !q.leftStack.IsEmpty() {
		q.rightStack.Push(q.leftStack.Pop())
	}

	return q.rightStack.Pop()
}

func (q *Queue) peek() int {
	if q.IsEmpty() {
		return int(-1)
	}

	if !q.rightStack.IsEmpty() {
		return q.rightStack.GetTop()
	}

	for !q.leftStack.IsEmpty() {
		q.rightStack.Push(q.leftStack.Pop())
	}

	return q.rightStack.GetTop()
}
