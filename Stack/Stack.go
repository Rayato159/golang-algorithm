package main

import "fmt"

type Stack struct {
	Data []int
	Size int
	Top  int
}

func NewStack(size int) *Stack {
	return &Stack{
		Data: make([]int, size),
		Size: size,
		Top:  -1,
	}
}

func (s *Stack) Push(x int) {
	if s.Top == s.Size-1 {
		fmt.Println("error, stack overflow")
	} else {
		s.Top++
		s.Data[s.Top] = x
	}
}

func (s *Stack) Pop() int {
	var x int
	if s.Top == -1 {
		fmt.Println("error, stack is empty")
	} else {
		x = s.Data[s.Top]
		s.Data[s.Top] = 0
		s.Top--
	}
	return x
}

func (s *Stack) Peek(i int) int {
	var x int
	if s.Data[i] == 0 {
		fmt.Println("error, not found.")
	} else {
		x = s.Data[i]
	}
	return x
}

func main() {
	s := NewStack(5)
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Push(50)
	fmt.Println(s.Data)
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println(s.Data)
}
