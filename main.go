package main

import (
	collection "DataStrictureDemo/collections"
	"fmt"
)

func main() {
	/*var stack Stack
	stack.Push(1)
	stack.Push("ABC")
	stack.Push(1.4)
	stack.Push("Lifting")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Dump())
	fmt.Println(stack.IsEmpty())
	items := stack.Dump()
	stack.Reset()
	fmt.Println(stack.items)
	fmt.Println(items)*/
	var stack collection.Stack
	stack.Push(2)
	var queue Queue
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	fmt.Println(queue.items)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.items)

}
