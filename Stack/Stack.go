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
	if s.Data[s.Top+i-1] <= 0 {
		fmt.Println("error, invalid position")
	} else {
		x = s.Data[s.Top+i-1]
	}
	return x
}

func main() {
	s := NewStack(5)
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s.Data, s.Top)
	fmt.Println(s.Peek(1))
}
