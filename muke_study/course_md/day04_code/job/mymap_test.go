package main

import (
	"fmt"
	"sync"
)
import "testing"

/*
选做：编写 benchmark
比较 MyMap 与 sync.Map 的同名函数性能差异(LoadAndDelete 可以不用比较，该函数在 Go 1.15 引入，我们的作业环境是 1.14.12)，
输出相应的性能报告(注意，你应该使用 RunParallel)，将性能比较结果输出为 markdown 文件
*/

func BenchmarkSyncMap_Store(b *testing.B) {
	var scene sync.Map
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			scene.Store("greece", 97)
		}
	})
}
func BenchmarkMMap_Store(b *testing.B) {
	var scene MMap
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			scene.Store("greece", 12)
		}
	})
}


func BenchmarkSyncMap_Load(b *testing.B) {
	var scene sync.Map
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			scene.Store("greece", "www")
			scene.Load("greece")
		}
	})
}

func BenchmarkMMap_Load(b *testing.B) {
	var scene MMap
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			scene.Store("greece", "www")
			scene.Load("greece")
		}
	})
}
func BenchmarkSyncMap_Delete(b *testing.B) {
	var scene sync.Map
	scene.Store("greece", "www")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			scene.Delete("greece")
		}
	})
}

func BenchmarkMMap_Delete(b *testing.B) {
	var scene MMap
	scene.Store("greece", "www")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			scene.Delete("greece")
		}
	})
}

func BenchmarkSyncMap_LoadOrStore(b *testing.B) {
	var scene sync.Map
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			scene.LoadOrStore("greece", "www")
		}
	})
}

func BenchmarkMMap_LoadOrStore(b *testing.B) {
	var scene MMap
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			scene.LoadOrStore("greece", "www")
		}
	})
}

func BenchmarkSyncMap_Range(b *testing.B) {
	var scene sync.Map
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// 遍历所有sync.Map中的键值对
			scene.Range(func(k, v interface{}) bool {
				fmt.Sprint("iterate:", k, v)
				return true
			})
		}
	})
}

func BenchmarkMMap_Range(b *testing.B) {
	var scene MMap
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// 遍历所有sync.Map中的键值对
			scene.Range(func(k, v interface{}) bool {
				fmt.Sprint("iterate:", k, v)
				return true
			})
		}
	})
}