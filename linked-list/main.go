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

	newList.AddBegining(1)
	newList.AddBegining(2)
	newList.AddEnd(3)
	newList.AddEnd(4)
	newList.AddEnd(5)
	newList.AddEnd(6)
	newList.AddEnd(7)
	newList.Forward()
	println("input 8 di posisi 6")
	newList.InsertSpesifik(6, 8)
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
}

func main() {
	//SingleList()
	DoubleList()
}
