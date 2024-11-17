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

	data.Arrays = data.Push(10)
	data.Arrays = data.Push(14)
	data.Arrays = data.Push(14)

	fmt.Println(data.Arrays)

	newData := arrays.Data{
		Arrays: []int{1, 2, 3, 2, 3, 4, 5, 6, 7, 8},
	}

	fmt.Println(newData.ReturnFirstRecurringv())
	//fmt.Println(data.DeleteSameValue())

	fmt.Println(newData.Sum(10))
}
