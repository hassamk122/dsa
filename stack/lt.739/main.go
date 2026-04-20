package main

import "fmt"

/*
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i]
is the number of days you have to wait after the
ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.
*/
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
		return 0
	}

	lastElement := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return lastElement
}

func (s *MonotonicStack) GetTop() int {
	if s.IsEmpty() {
		return 0
	}

	return s.arr[len(s.arr)-1]
}

func main() {
	arr := []int{30, 40, 50, 60}

	result := dailyTemperatures(arr)

	fmt.Println(result)

}

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 || len(temperatures) == 1 {
		return temperatures
	}

	stack := NewMonotonicStack()
	resultArr := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {

		for temperatures[i] >= temperatures[stack.GetTop()] && !stack.IsEmpty() {
			stack.Pop()
		}

		if stack.GetTop() == 0 {
			resultArr[i] = stack.GetTop()
		} else {
			resultArr[i] = stack.GetTop() - i
		}

		stack.Push(i)
	}

	return resultArr
}
