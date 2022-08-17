package main

import "fmt"

type Queue struct {
	Data  []int
	First int
	Rear  int
}

func NewQueue(size int) *Queue {
	return &Queue{
		Data:  make([]int, size+1),
		First: 0,
		Rear:  0,
	}
}

func (q *Queue) Enqueue(x int) {
	if (q.First+1)%(cap(q.Data)-1) == q.Rear {
		fmt.Println("error, queue is full.")
	} else {
		q.Rear = (q.Rear + 1) % (cap(q.Data) - 1)
		q.Data[q.Rear] = x
		fmt.Println(q.Data, q.First, q.Rear)
	}
}

func (q *Queue) Dequeue() {
	if q.First == q.Rear {
		fmt.Println("error, queue is empty.")
	} else {
		q.First = (q.First + 1) % (cap(q.Data) - 1)
		q.Data[q.First] = 0
		fmt.Println(q.Data, q.First, q.Rear)
	}
}

func main() {
	q := NewQueue(5)
	q.Enqueue(1)
	q.Dequeue()
	q.Dequeue()
}
