package stacks

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Push(v int) {
	newNode := &Node{Data: v}

	if list.Head != nil {
		newNode.Next = list.Head
		list.Head = newNode
	} else {
		list.Head = newNode
	}
}

func (list *LinkedList) Pop() {
	if list.Head == nil {
		return
	}

	if list.Head != nil {
		list.Head = list.Head.Next
		return
	}
}

func (list *LinkedList) Lookup() {
	if list.Head == nil {
		return
	}

	current := list.Head

	for current != nil {
		fmt.Printf("%d \n", current.Data)
		current = current.Next
	}
}

func (list *LinkedList) Peek() {
	fmt.Printf("%d \n", list.Head.Data)
}

func (list *LinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *LinkedList) Size() int {
	result := 0
	current := list.Head
	for current != nil {
		current = current.Next
		result++
	}

	return result
}

func (list *LinkedList) Search(v int) int {
	position := 1
	current := list.Head

	for current != nil && current.Data != v {
		current = current.Next
		position++
	}

	if current == nil {
		return -1
	}

	return position
}

func (list *LinkedList) DuplicateTop() {
	if list.Head == nil {
		return
	}

	current := list.Head
	topValue := current.Data
	newNode := &Node{Data: topValue}

	if list.Head != nil {
		newNode.Next = list.Head
		list.Head = newNode
	}
}

func (list *LinkedList) Clear() {
	list.Head = nil
}
