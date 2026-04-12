package main

import "fmt"

type Stack struct {
	arr []int
}

func (s *Stack) GetTop() int {
	return s.arr[len(s.arr)-1]
}

func (s *Stack) Push(data int) {
	s.arr = append(s.arr, data)
}

func (s *Stack) Pop() {
	s.arr = s.arr[:len(s.arr)-1]
}

func main() {
	s := &Stack{}
	s.Push(1)
	s.Push(2)

	s.Push(3)
	s.Pop()

	fmt.Println(s.GetTop())
	fmt.Println(s.arr)
}
