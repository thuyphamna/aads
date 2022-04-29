package main

import (
	"fmt"
	"stacksandqueues/queues"
)

func main() {
	// s := stacks.Stack{}
	// s.Push(1)
	// s.Push(2)
	// s.Push(3)
	// s.Push(4)
	// s.Push(5)
	// s.Push(5)
	// s.Push(5)
	// s.Push(5)

	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(len(s))

	q := queues.Queue{}
	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(1)

	fmt.Println(q.Dequeue())
	fmt.Println(len(q))

}
