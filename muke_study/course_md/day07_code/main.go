package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go deferCall()
	var i = 0
	for i = 0; i < 10; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Minute)
}

func deferCall() {
	defer func() {fmt.Println("打印前")}()
	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()

	defer func() {
		if err := recover();err!=nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		defer func() {
			panic("panic again and again") // --
		}()
		panic("panic again")
	}()
	panic("触发异常")
}
