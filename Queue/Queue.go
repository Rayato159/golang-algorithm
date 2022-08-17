package main

import "fmt"

type Queue struct {
	Data  []int
	Size  int
	Front int
	Rear  int
}

func NewQueue(size int) *Queue {
	return &Queue{
		Data:  make([]int, size+2),
		Size:  size + 1,
		Front: 0,
		Rear:  0,
	}
}

func (q *Queue) Enqueue(x int) {
	if (q.Rear+1)%q.Size == q.Front {
		fmt.Println("error, queue is full")
	} else {
		q.Rear = (q.Rear + 1) % q.Size
		q.Data[q.Rear] = x
		fmt.Println(q.Data, q.Front, q.Rear)
	}
}

func (q *Queue) Dequeue() int {
	var x int
	if q.Front == q.Rear {
		fmt.Println("error, queue is empty")
	} else {
		q.Front = (q.Front + 1) % q.Size
		x = q.Data[q.Front]
		q.Data[q.Front] = 0
		fmt.Println(q.Data, q.Front, q.Rear)
	}
	return x
}

func main() {
	q := NewQueue(5)
	fmt.Println(q)
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
}
