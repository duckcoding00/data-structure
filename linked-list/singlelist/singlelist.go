package singlelist

import "fmt"

type Node struct {
	Next *Node // references to next node
	Data int
}

type SingleList struct {
	Head *Node // first node are head
}

func (list *SingleList) Add(v int) {
	// initiate new node
	newNode := Node{Data: v}

	// checking jika list head tidak kosong atau nil
	if list.Head != nil {
		newNode.Next = list.Head
		list.Head = &newNode
	} else {
		list.Head = &newNode
	}
}

func (list *SingleList) Remove(v int) {
	// check jika list head bernilai kosong atau nil
	if list.Head == nil {
		return
	}

	// check jika data berada di list head
	if list.Head.Data == v {
		// maka head akan dipindah ke node selanjutnya
		list.Head = list.Head.Next
	}

	// jika data berada di rentang list
	previousToDelete := list.Head
	for previousToDelete.Next != nil && previousToDelete.Next.Data != v {
		previousToDelete = previousToDelete.Next
	}

	// jika node selanjutnya sama nil
	if previousToDelete.Next == nil {
		return
	}

	previousToDelete.Next = previousToDelete.Next.Next
}

func (list *SingleList) Print() {
	if list.Head == nil {
		return
	}

	current := list.Head
	for current != nil {
		fmt.Printf("%d\n", current.Data)
		current = current.Next
	}
}
