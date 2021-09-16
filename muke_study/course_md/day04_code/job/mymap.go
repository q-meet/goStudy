package main

import (
	"fmt"
	"sync"
)

// 只支持 int 即可。
/*
必做：
封装一个数据结构 MyMap，实现并发安全的 Load，Store，Delete，LoadAndDelete，LoadOrStore 几个 API(禁止使用 sync.Map)，不用考虑性能，文件在https://golearn.coding.net/p/gonggongbanji/files/all/DF18 中的 mymap.go

*/
type MMap struct {
	mux  sync.RWMutex
	elem map[interface{}]interface{}
	init int
}
func (m *MMap) Init(){
	if m.init == 1 {
		return
	}
	m.init = 1
	m.elem = make(map[interface{}]interface{})
}

func (m *MMap) Load(key interface{}) (value interface{}, ok bool) {
	m.mux.RLock()
	value, ok = m.elem[key]
	m.mux.RUnlock()
	return
}

func (m *MMap) Store(key, value interface{}) {
	m.Init()
	m.mux.Lock()
	m.elem[key] = value
	m.mux.Unlock()
}

func (m *MMap) Delete(key interface{}) {
	m.LoadAndDelete(key)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *MMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	m.mux.Lock()
	m.Init()
	if actual, loaded = m.elem[key]; loaded {
		m.mux.Unlock()
		return
	}
	m.elem[key] = value
	m.mux.Unlock()
	return value, false
}

func (m *MMap) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	m.mux.Lock()
	if value, loaded = m.elem[key]; loaded {
		delete(m.elem, key)
	}
	m.mux.Unlock()
	return
}

func (m *MMap) Range(f func(key, value interface{}) bool) {
	m.mux.RLock()
	for k := range m.elem {
		v, ok := m.elem[k]
		if !ok {
			continue
		}
		if !f(k, v) {
			break
		}
	}
	m.mux.RUnlock()
}


func main() {
	var scene MMap
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	fmt.Println(scene.Load("greece"))
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	scene.LoadOrStore("london","哈哈哈")
	fmt.Println(scene.Load("london"))
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true

	})
	syncMap()
}

func syncMap() {

	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	fmt.Println(scene.Load("greece"))
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	scene.LoadOrStore("london","哈哈哈")
	fmt.Println(scene.Load("london"))
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
