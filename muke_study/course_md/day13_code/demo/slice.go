package main

import "fmt"

type MySlice []int

func (m MySlice) append(i int) MySlice {
	return append(m, i)
}

func main() {
	var m = make(MySlice, 1, 1)
	fmt.Println(m)
	numbers := append(m, 1)

	fmt.Println(numbers)

	fmt.Println(m)
}
