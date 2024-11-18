package main

import (
	"fmt"
	"linked-list/singlelist"
)

func main() {
	// create new singlelist

	newSingleList := singlelist.SingleList{}

	newSingleList.Add(2)
	newSingleList.Add(4)
	newSingleList.Add(5)
	newSingleList.Add(1)
	newSingleList.Add(8)
	newSingleList.Print()

	fmt.Println("")

	newSingleList.Remove(5)
	newSingleList.Print()
}
