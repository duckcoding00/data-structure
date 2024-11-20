package queues

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Data int
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (list *LinkedList) Enqueue(v int) {
	newNode := &Node{Data: v}

	if list.Head == nil && list.Tail == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
	}
}

func (list *LinkedList) Traverse() {
	if list.Tail == nil {
		return
	}

	current := list.Tail
	for current != nil {
		fmt.Printf("%d \n", current.Data)
		current = current.Prev
	}
}

func (list *LinkedList) Dequeue() {
	if list.Tail == nil {
		return
	}

	if list.Tail.Prev == nil {
		list.Head = nil
		list.Tail = nil
		return
	}

	if list.Tail != nil {
		list.Tail = list.Tail.Prev
		list.Tail.Next = nil
	}
}

func (list *LinkedList) Peek() int {
	return list.Tail.Data
}

func (list *LinkedList) Last() int {
	return list.Head.Data
}

func (list *LinkedList) IsEmpty() bool {
	return list.Tail == nil
}

func (list *LinkedList) Size() int {
	result := 0
	current := list.Tail

	for current != nil {
		current = current.Prev
		result++
	}

	return result
}

func (list *LinkedList) Search(v int) int {
	position := 1
	current := list.Tail

	for current != nil && current.Data != v {
		current = current.Prev
		position++
	}

	if current == nil {
		return -1
	}

	return position
}
