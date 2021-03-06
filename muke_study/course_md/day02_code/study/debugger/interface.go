package main

import (
	"fmt"
	"io"
	"os"
)

var (
	v  interface{}
	r  io.Reader
	f  *os.File
	fn os.File
)

func main() {
	fmt.Println(v == nil)
	fmt.Println(r == nil)
	fmt.Println(f == nil)
	v = r
	fmt.Println(v == nil)
	v = fn
	fmt.Println(v == nil)
	v = f
	fmt.Println(v == nil)
	r = f
	fmt.Println(r == nil)
}