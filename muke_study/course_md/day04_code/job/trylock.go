package main

import (
	"fmt"
)

/*
trylock顾名思义，尝试加锁，加锁成功执行后续流程，如果加锁失败的话也不会阻塞，而会直接返回加锁的结果。
在Go语言中我们可以用大小为1的Channel来模拟trylock：
*/
type MyLock struct {
	// you should have a channel here
	ch chan struct{}
}

func NewMutex() *MyLock {
	mu := &MyLock{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

func (m *MyLock) Lock() {
	<-m.ch
}

func (m *MyLock) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (m * MyLock) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

func (m *MyLock) IsLocked() bool {
	return len(m.ch) == 0
}
func main() {
	m := NewMutex()
	fmt.Println( len(m.ch)) //1
	fmt.Printf("TryLock: %t\n", m.TryLock()) //true
	fmt.Println( len(m.ch)) //0
	m.Unlock()
	fmt.Println( len(m.ch)) //1
	fmt.Printf("TryLock: %t\n", m.TryLock()) //true
	fmt.Printf("TryLock: %t\n", m.TryLock()) //false
	m.Unlock()


}

