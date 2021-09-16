package main

import (
	"fmt"
)

func main() {
	var r rune = -1
	fmt.Println(len(string(r)))


	var aa int32 =  1111111
	var bb int32 =  2147483647
	fmt.Println(len(string(aa)))
	fmt.Println("bb", len(string(bb)))

}