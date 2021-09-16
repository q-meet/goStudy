package main

import (
	"fmt"
	"sync"
)

// 只支持 int 即可。

type MyMap struct {
}

func (m *MyMap) Load(key interface{}) (value interface{}, ok bool) {
	return nil, false
}

func (m *MyMap) Store(key, value interface{}) {
}

func (m *MyMap) Delete(key interface{}) {

}

func (m *MyMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	return nil, false
}

func (m *MyMap) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	return nil, false
}

func main() {

	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
