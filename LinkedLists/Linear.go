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
	fmt.Println(head, last)
	last = last.AddNode(20)
	fmt.Println(head, last)
	last = last.AddNode(30)
	fmt.Println(head, last)
	last = last.AddNode(40)
	fmt.Println(head, last)
	last = last.AddNode(50)
	fmt.Println(head, last)
	head.Display()
}
