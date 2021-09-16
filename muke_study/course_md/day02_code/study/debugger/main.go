package main

import (
	"bytes"
	"fmt"
	"io"
)

var (
	a *bytes.Buffer = nil
	b io.Writer
)

func set(v *bytes.Buffer) {
	if v == nil {
		fmt.Println("v is nil")
	}
	b = v
}

func get() {
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil")
	}
}

func main() {
	set(nil)
	get()


}
