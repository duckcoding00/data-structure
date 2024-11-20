package main

import (
	"fmt"
	"linked-list/doublelist"
	"linked-list/singlelist"
)

func SingleList() {
	newList := singlelist.LinkedList{}

	newList.Add(1)
	newList.Add(2)
	newList.Add(3)
	newList.Add(6)
	newList.Add(4)
	newList.Add(5)
	newList.Print()
	fmt.Printf("\ndelete first data\n")
	newList.DeleteFirst()
	newList.Print()
	fmt.Printf("\ndelete last data\n")
	newList.DeleteLast()
	newList.Print()
	fmt.Printf("\ndelete spesific value\n")
	newList.Remove(4)
	newList.Print()
}

func DoubleList() {
	newList := doublelist.LinkedList{}

	AddBegining := []int{1, 2, 3, 4, 5}
	for _, v := range AddBegining {
		newList.AddBegining(v)
	}
	newList.Forward()

	AddEnd := []int{7, 8, 9}
	for _, v := range AddEnd {
		newList.AddEnd(v)
	}
	newList.Forward()

	newList.InsertSpesifik(1, 8)
	fmt.Printf("input %d di posisi %d \n", 1, 8)
	newList.Forward()
	println("delete first")
	newList.DeleteFirst()
	newList.Forward()
	println("delete last")
	newList.DeleteLast()
	newList.Forward()
	println("delete 4")
	newList.DeleteSpesifik(8)
	newList.Forward()
	println("reverse")
	newList.Reverse()

	newList.Sorting()
	newList.Forward()
}

func main() {
	//SingleList()
	DoubleList()
}
