package main

import (
	"fmt"
	"stack_queue/structures"
)

func main() {
	myStack := structures.Stack{}
	myQueue := structures.Queue{}
	myStack.Push("3")
	myStack.Push("2")
	myStack.Push("1")
	fmt.Println(myStack.IsEmpty())
	fmt.Println(myStack.Peek())
	myQueue.Enqueue("1")
	myQueue.Enqueue("2")
	myQueue.Enqueue("3")
	fmt.Println(myQueue.IsEmpty())
	fmt.Println(myQueue.Peek())
}
