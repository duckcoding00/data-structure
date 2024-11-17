package linkedlist

import "fmt"

type node struct {
	next *node // reference next node
	Data int
}

type LinkedList struct {
	head   *node // first node
	length int
}

type Linkedlist interface {
	Prepend(v int)
	Print()
	Delete(v int)
}

func (l *LinkedList) Prepend(v int) {
	newNode := node{Data: v}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
	} else {
		l.head = &newNode
	}
	l.length++
}

func (l *LinkedList) Print() {
	if l.head == nil {
		return
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.next
	}
}

func (l *LinkedList) Delete(v int) {
	if l.length == 0 {
		return
	}

	if l.head.Data == v {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next != nil && previousToDelete.next.Data != v {
		previousToDelete = previousToDelete.next
	}

	if previousToDelete.next == nil {
		return
	}

	previousToDelete.next = previousToDelete.next.next
	l.length--
}
