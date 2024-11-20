package main

import (
	"fmt"
	"stacks-queues/queues"
	"stacks-queues/stacks"
)

func Stack() {
	newStack := stacks.LinkedList{}

	fmt.Println("PUSH 1->2->3->4->5")
	data := []int{1, 2, 3, 4, 5}
	for _, v := range data {
		newStack.Push(v)
	}

	fmt.Println("Traverse")
	newStack.Traverse()
	fmt.Println("PEEK")
	newStack.Peek()
	fmt.Println("")
	fmt.Println("POP last item")
	newStack.Pop()
	fmt.Println("Traverse")
	newStack.Traverse()
	fmt.Println("PEEK")
	newStack.Peek()

	fmt.Println(newStack.IsEmpty())
	fmt.Println(newStack.Size())
	fmt.Println(newStack.Search(2))

	newStack.DuplicateTop()
	fmt.Println("Traverse")
	newStack.Traverse()

	newStack.Clear()
	fmt.Println("Traverse")
	newStack.Traverse()
	fmt.Println("\nPUSH 1->2->3->4->5")
	for _, v := range data {
		newStack.Push(v)
	}

	fmt.Println("Traverse")
	newStack.Traverse()
}

func Queue() {
	newQueue := queues.LinkedList{}

	data := []int{1, 2, 3, 4, 5, 6}

	for _, v := range data {
		newQueue.Enqueue(v)
	}
	newQueue.Traverse()
	println("")
	newQueue.Dequeue()
	newQueue.Traverse()
	println("")
	fmt.Println(newQueue.Peek())
	fmt.Println(newQueue.IsEmpty())
	fmt.Println(newQueue.Size())
	fmt.Println(newQueue.Search(4))
}

func main() {
	//Stack()
	Queue()
}
