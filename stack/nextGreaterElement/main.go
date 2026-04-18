package main

import "fmt"

type MonotonicStack struct {
	arr []int
}

func NewMonotonicStack() MonotonicStack {
	return MonotonicStack{
		arr: make([]int, 0),
	}
}

func (s *MonotonicStack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *MonotonicStack) Push(data int) {

	s.arr = append(s.arr, data)
}

func (s *MonotonicStack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	lastElement := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return lastElement
}

func (s *MonotonicStack) GetTop() int {
	if s.IsEmpty() {
		return -1
	}

	return s.arr[len(s.arr)-1]
}

func main() {
	arr := []int{3, 1, 5, 2, 1, 0, 7, 9, 2, 3}

	result := NextGreaterElement(arr)

	fmt.Println(result)
}

func NextGreaterElement(arr []int) []int {
	stack := NewMonotonicStack()
	resultArr := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {

		for arr[i] > stack.GetTop() && !stack.IsEmpty() {
			stack.Pop()
		}

		resultArr[i] = stack.GetTop()

		stack.Push(arr[i])
	}
	return resultArr
}
