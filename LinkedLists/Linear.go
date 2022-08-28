package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

func (n *Node) AddNode(data int) *Node {
	t := &Node{
		Data: data,
		Next: nil,
	}
	n.Next = t
	n = t
	return t
}

func (n *Node) Display() {
	if n != nil {
		fmt.Println(n.Data)
		n.Next.Display()
	}
}

func main() {
	head := NewNode(10)
	last := head
	last = last.AddNode(20)
	last = last.AddNode(30)
	last = last.AddNode(40)
	last = last.AddNode(50)
	head.Display()
}
