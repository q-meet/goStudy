package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
有三个函数 分别可以打印 “cat" "dog" "fish"
要求每个函数都起一个goroutine 按照 cat dog fish的顺序打印在屏幕上 100次
*/

var mu sync.WaitGroup

var Cat int32 = 1
var Dog int32 = 1
var Hish int32 = 1

func main() {
	readCat := make(chan struct{}, 1)
	readDog := make(chan struct{}, 0)
	readHish := make(chan struct{}, 0)
	readCat <- struct{}{}
	for i := 0; i < 100; i++ {
		mu.Add(3)
		go fmts(readCat, readDog, "cat", &Cat)
		go fmts(readDog, readHish, "Dog", &Dog)
		go fmts(readHish, readCat, "Hish", &Hish)
		//go fmtCat(readCat, readDog)
		//go fmtDog(readDog, readHish)
		//go fmtHish(readHish, readCat)
	}
	mu.Wait()
}

func fmtCat(read, w chan struct{}) {
	<-read
	fmt.Println("Cat", Cat)
	atomic.AddInt32(&Cat, 1)
	w <- struct{}{}
	mu.Done()
}

func fmtDog(read, w chan struct{}) {
	<-read
	fmt.Println("Dog", Dog)
	atomic.AddInt32(&Dog, 1)
	w <- struct{}{}
	mu.Done()
}

func fmtHish(read <-  chan struct{} , w chan <-  struct{}) {
	<-read
	fmt.Println("Hish", Hish)
	atomic.AddInt32(&Hish, 1)
	w <- struct{}{}
	mu.Done()
}

func fmts(read <-  chan struct{} , w chan <-  struct{}, content string, num *int32) {
	<-read
	fmt.Println(content, *num)
	atomic.AddInt32(num, 1)
	w <- struct{}{}
	mu.Done()
}
