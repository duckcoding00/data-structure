package main

import (
	"linked-list/linkedlist"
)

func main() {
	newLinkedList := linkedlist.LinkedList{}

	newLinkedList.Prepend(1)
	newLinkedList.Prepend(4)
	newLinkedList.Prepend(5)
	newLinkedList.Prepend(10)
	newLinkedList.Prepend(24) // last input are head

	newLinkedList.Print()

	newLinkedList.Delete(5)

	newLinkedList.Print()
}
