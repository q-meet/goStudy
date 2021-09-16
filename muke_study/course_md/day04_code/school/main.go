package main

import "sync"

// 只支持 int 即可。

type MyMap struct {
	rwLock  sync.RWMutex
	realMap map[interface{}]interface{}
}

func (m *MyMap) Load(key interface{}) (value interface{}, ok bool) {
	m.rwLock.RLock()
	if value, ok = m.realMap[key]; ok {
		m.rwLock.RUnlock()
		return
	}
	m.rwLock.RUnlock()
	return nil, false
}

func (m *MyMap) Store(key, value interface{}) {
	m.rwLock.Lock()
	m.realMap[key] = value
	m.rwLock.Unlock()
}

func (m *MyMap) Delete(key interface{}) {
	m.rwLock.Lock()
	delete(m.realMap, key)
	m.rwLock.Unlock()
}

func (m *MyMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	m.rwLock.Lock()
	if actual, loaded = m.realMap[key]; loaded {
		m.rwLock.Unlock()
		return
	}
	m.realMap[key] = value
	m.rwLock.Unlock()
	return value, false
}

func (m *MyMap) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	m.rwLock.Lock()
	if value, loaded = m.realMap[key]; loaded {
		delete(m.realMap, key)
		m.rwLock.Unlock()
		return
	}
	m.rwLock.Unlock()
	return nil, false
}

func main() {

	var scene MyMap
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
}