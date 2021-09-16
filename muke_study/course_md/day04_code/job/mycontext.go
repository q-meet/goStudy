package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
在 context 学习过程中，我们知道，在子节点中 WithValue 的数据，父节点是查不到的。
请实现一个 MyContext，其 WithValue 方法赋值的 k，v 在父节点中也可以查得到。
*/
type MyContext struct {
	ch chan int
	elem map[interface{}]interface{}
}
//定义string类型
type TraceCode string

func (m *MyContext) WithCancel(ctx context.Context) (*MyContext, context.Context, func()){
	m.ch = make(chan int)
	return m, ctx, func() {
		m.ch <- 1
	}
}

func (m * MyContext) WithValue(parent context.Context, key, val interface{}) context.Context {
	m.elem = make(map[interface{}]interface{})
	m.elem[key] = val
	return parent
}


func (m * MyContext) Value(key interface{}) (interface{}, bool) {
	value, ok := m.elem[key]
	return value, ok
}

var wgV sync.WaitGroup

func workerV(m *MyContext, ctx context.Context) {
	ctx = m.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	fmt.Println(m.elem)
LOOP:
	for {
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <- m.ch:
			break LOOP
		default:
		}
	}
	fmt.Println("workerV done!")
	wgV.Done()
}
func main() {
	var MyCon MyContext
	m, ctx, cancel := MyCon.WithCancel(context.Background())

	wgV.Add(1)

	go workerV(m, ctx)
	time.Sleep(time.Second * 1)

	traceCode, ok := m.Value(TraceCode("TRACE_CODE")) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
	fmt.Printf("workerV, trace code:%s\n", traceCode)

	cancel() // 当我们取完需要的整数后调用cancel
	wgV.Wait()

	fmt.Println("over")
}
