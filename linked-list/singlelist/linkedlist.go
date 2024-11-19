package singlelist

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Print() {
	if list.Head == nil {
		return
	}

	current := list.Head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
}

func (list *LinkedList) Add(v int) {
	newNode := Node{Data: v}

	if list.Head != nil {
		newNode.Next = list.Head
		list.Head = &newNode
	} else {
		list.Head = &newNode
	}
}

func (list *LinkedList) DeleteFirst() {
	if list.Head == nil {
		return
	}

	if list.Head.Next == nil {
		list.Head = nil
		return
	}

	if list.Head.Next != nil {
		list.Head = list.Head.Next
	}
}

func (list *LinkedList) DeleteLast() {
	if list.Head == nil {
		return
	}

	if list.Head.Next == nil {
		list.Head = nil
		return
	}

	current := list.Head
	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
}

func (list *LinkedList) Remove(v int) {
	if list.Head == nil {
		return
	}

	if list.Head.Data == v {
		list.Head = list.Head.Next
	}

	previousToDelete := list.Head
	for previousToDelete.Next != nil && previousToDelete.Next.Data != v {
		previousToDelete = previousToDelete.Next
	}

	if previousToDelete.Next == nil {
		return
	}

	previousToDelete.Next = previousToDelete.Next.Next
}
