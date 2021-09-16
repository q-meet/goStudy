package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(1000)
	a := 0
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				a += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(a)
}
