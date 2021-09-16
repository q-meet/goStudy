package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	GMP:
	M(线程),任务消费者; G(goroutine),计算任务; P,可以使用的CPU的token

	队列:
	goroutine处理有三个队列(多级队列减少锁竞争)

	P的本地runnext字段(只能存储一个)

	P的local run queue(数组固定大小256)

	global run queue(切片 无限大小)

	调度循环：
	线程M在持有P的情况下不断消费运行队列中的G的过程

	处理阻塞:
	1.可以接管的阻塞：channel收发, 加锁, 网络连接读/写, select
	2.不可接管的阻塞: syscall, cgo, 长时间运行需要剥离P执行
*/
func main() {

	// 要给无聊的顺序输出的问题
	// 定义使用一个M
	runtime.GOMAXPROCS(1)

	for i := 0; i < 270; i++ {
		i := i
		go func() {
			fmt.Println("A:", i)
		}()
	}

	time.Sleep(time.Hour)
	/*
		var ch = make(chan int)
		<-ch*/
}
