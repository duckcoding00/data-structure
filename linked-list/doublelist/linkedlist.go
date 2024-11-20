package doublelist

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

func (list *LinkedList) Forward() {
	if list.Head == nil {
		return
	}

	current := list.Head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println("")
}

func (list *LinkedList) Reverse() {
	if list.Tail == nil {
		return
	}

	current := list.Tail
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Prev
	}
	fmt.Println("")
}

func (list *LinkedList) AddBegining(v int) {
	newNode := &Node{Data: v}

	if list.Head != nil {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
	} else {
		list.Head = newNode
		list.Tail = newNode
	}
}

func (list *LinkedList) AddEnd(v int) {
	newNode := &Node{Data: v}

	if list.Tail != nil {
		newNode.Prev = list.Tail
		list.Tail.Next = newNode
		list.Tail = newNode
	} else {
		list.Head = newNode
		list.Tail = newNode
	}
}

func (list *LinkedList) InsertSpesifik(p, v int) {
	if p < 1 {
		fmt.Println("invalid position")
		return
	}

	newNode := Node{Data: v}
	current := list.Head
	position := 1

	for current != nil && position < p-1 {
		current = current.Next
		position++
	}

	if current == nil {
		fmt.Println("position out of range")
		return
	}

	newNode.Next = current.Next
	newNode.Prev = current

	if current.Next != nil {
		current.Next.Prev = &newNode
	} else {
		list.Tail = &newNode
	}

	current.Next = &newNode
}

func (list *LinkedList) DeleteFirst() {
	if list.Head == nil {
		return
	}

	if list.Head.Next == nil {
		list.Head = nil
		list.Tail = nil
		return
	}

	if list.Head != nil {
		list.Head = list.Head.Next
		list.Head.Prev = nil
	}
}

func (list *LinkedList) DeleteLast() {
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

func (list *LinkedList) DeleteSpesifik(v int) {
	if list.Head == nil {
		return
	}

	if list.Head.Data == v {
		if list.Head == list.Tail {
			list.Head = nil
			list.Tail = nil
		}
		list.Head = list.Head.Next
		list.Head.Prev = nil
		return
	}

	if list.Tail.Data == v {
		list.Tail = list.Tail.Prev
		list.Tail.Next = nil
		return
	}

	current := list.Head
	for current != nil && current.Data != v {
		current = current.Next
	}

	if current == nil {
		return
	}

	current.Prev.Next = current.Next
	if current.Next != nil {
		current.Next.Prev = current.Prev
	}
}

func (list *LinkedList) Sorting() {
	if list.Head == nil || list.Tail == nil {
		return
	}

	swapped := true

	for swapped {
		current := list.Head
		swapped = false

		for current.Next != nil {
			if current.Data > current.Next.Data {
				temp := current.Data
				current.Data = current.Next.Data
				current.Next.Data = temp

				swapped = true
			}

			current = current.Next
		}
	}
}
