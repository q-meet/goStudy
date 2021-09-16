package main

import "fmt"

func main() {
	d := w()

	if r := d(-1); r != nil {
		fmt.Println("Good result:", r)
	}
}

func w() func(arg int) interface{} {
	return func(arg int) interface{} {
		var r *struct{} = nil
		if arg > 0 {
			r = &struct{}{}
		} else {
			return nil
		}
		return r
	}
}
