package arrays

import "fmt"

type Data struct {
	Arrays []int
}

type Operation interface {
	Push(v int) []int
	Delete(t int) []int
	Shift(v int) []int
	Unshift(i, v int) []int
}

func (d *Data) Push(v int) []int {
	newSlice := make([]int, len(d.Arrays)+1)

	for i := range d.Arrays {
		newSlice[i] = d.Arrays[i]
	}

	newSlice[len(d.Arrays)] = v

	return newSlice
}

func (d *Data) Delete(t int) []int {
	newSlice := make([]int, 0, len(d.Arrays)-1)

	for i, v := range d.Arrays {
		if v == t {
			newSlice = append(newSlice, d.Arrays[:i]...)
			newSlice = append(newSlice, d.Arrays[i+1:]...)
		} else {
			fmt.Println("target not in arrays")
		}
	}

	if len(newSlice) == 0 {
		return d.Arrays
	}

	return newSlice
}

func (d *Data) Shift(v int) []int {
	newSlice := make([]int, len(d.Arrays)+1)

	newSlice[0] = v

	for i := range d.Arrays {
		newSlice[i+1] = d.Arrays[i]
	}

	return newSlice
}

func (d *Data) Unshift(i, v int) []int {
	if i < 0 || i > len(d.Arrays) {
		fmt.Println("index not valid")
		return nil
	}

	newSlice := make([]int, len(d.Arrays)+1)

	for index := 0; index < i; index++ {
		newSlice[index] = d.Arrays[index]
	}

	newSlice[i] = v

	for index := i; index < len(d.Arrays); index++ {
		newSlice[index+1] = d.Arrays[index]
	}

	return newSlice
}
