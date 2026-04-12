package main

import "fmt"

/*
	Valid parenthesis
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.

	Input: s = "()[]{}"

Output: true
*/

type Stack struct {
	arr []rune
}

func (s *Stack) GetTop() rune {
	if s.IsEmpty() {
		return rune(' ')
	}
	return s.arr[len(s.arr)-1]
}

func (s *Stack) Push(data rune) {
	s.arr = append(s.arr, data)
}

func (s *Stack) Pop() {
	s.arr = s.arr[:len(s.arr)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func main() {

	s1 := "()"
	fmt.Println(s1, isValid(s1))

	s2 := "()[]{}"
	fmt.Println(s2, isValid(s2))

	s3 := "(]"
	fmt.Println(s3, isValid(s3))

	s4 := "([)]"
	fmt.Println(s4, isValid(s4))

	s5 := "]"
	fmt.Println(s5, isValid(s5))

}

func isValid(s string) bool {
	stack := &Stack{}

	for _, v := range s {

		switch v {
		case '(', '{', '[':
			stack.Push(v)
		case ')':
			if stack.GetTop() != '(' {
				return false
			}
			stack.Pop()
		case '}':
			if stack.GetTop() != '{' {
				return false
			}
			stack.Pop()
		case ']':
			if stack.GetTop() != '[' {
				return false
			}
			stack.Pop()
		}

	}

	if !stack.IsEmpty() {
		return false
	}

	return true
}
