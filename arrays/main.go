package main

import (
	"arrays/arrays"
	"fmt"
)

func main() {
	data := arrays.Data{
		Arrays: []int{},
	}

	data.Arrays = data.Push(1)
	data.Arrays = data.Push(5)
	data.Arrays = data.Push(10)
	fmt.Println(data.Arrays)

	data.Arrays = data.Delete(10)
	fmt.Println(data.Arrays)

	data.Arrays = data.Shift(100)
	fmt.Println(data.Arrays)

	data.Arrays = data.Unshift(2, 4)
	fmt.Println(data.Arrays)
}
