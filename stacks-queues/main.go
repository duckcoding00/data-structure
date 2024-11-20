package main

import (
	"fmt"
	"stacks-queues/stacks"
)

func Stack() {
	newStack := stacks.LinkedList{}

	fmt.Println("PUSH 1->2->3->4->5")
	data := []int{1, 2, 3, 4, 5}
	for _, v := range data {
		newStack.Push(v)
	}

	fmt.Println("LOOKUP")
	newStack.Lookup()
	fmt.Println("PEEK")
	newStack.Peek()
	fmt.Println("")
	fmt.Println("POP last item")
	newStack.Pop()
	fmt.Println("LOOKUP")
	newStack.Lookup()
	fmt.Println("PEEK")
	newStack.Peek()

	fmt.Println(newStack.IsEmpty())
	fmt.Println(newStack.Size())
	fmt.Println(newStack.Search(2))

	newStack.DuplicateTop()
	fmt.Println("LOOKUP")
	newStack.Lookup()

	newStack.Clear()
	fmt.Println("LOOKUP")
	newStack.Lookup()
	fmt.Println("\nPUSH 1->2->3->4->5")
	for _, v := range data {
		newStack.Push(v)
	}

	fmt.Println("LOOKUP")
	newStack.Lookup()
}

func main() {
	Stack()
}
