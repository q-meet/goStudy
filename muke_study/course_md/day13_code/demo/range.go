package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {
	var m = map[int]person{
		1: {
			age:  11,
			name: "abc",
		},
		2: {
			age:  22,
			name: "ddd",
		},
	}

	fmt.Println(m)

	var mm = map[int]*person{}
	for k, v := range m {
		w := v
		mm[k] = &w
	}
	fmt.Println(mm[1], mm[2])
}
