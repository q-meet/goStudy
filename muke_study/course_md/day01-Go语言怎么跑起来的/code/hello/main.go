package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	/*
	for {
		var x *int32
		*x = 0
	}
	*/
	// NumCPU 返回当前进程可以用到的逻辑核心数
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(4)
	// NumCPU 返回当前进程可以用到的逻辑核心数
	fmt.Println(runtime.NumCPU())
	fmt.Println("hello")
	for i := 0; i < 11; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Hour)
	//runtime.rt0_go()
	//runtime.runqput()
	//runtime.runqget()
	//runtime.globrunqput()
	//runtime.globrunqget()

	//runtime.schedule()
	//runtime.findrunnable()
	//runtime.sysmon()
}
