package main

import (
	"./Queue"
	"fmt"
)

func main() {
	//Array.AddFirst(6)
	//Array.AddFirst(2)
	//Array.AddLast(3)
	//Array.Del(1)
	//Array.Add(2, 5)
	//Array.GetRes()

	//var i Stack.Stack
	//s := Stack.Stacks{}
	//i = s
	//for k := 0; k < 6; k++ {
	//	i.Push(k)
	//}
	//i.Pop()

	queue := Queue.Queues{}
	for k := 0; k < 10; k++ {
		queue.EnQueue(k)
		if k%3 == 2 {
			queue.DeQueue()
		}
	}

	fmt.Println(queue.GetArr())
}
