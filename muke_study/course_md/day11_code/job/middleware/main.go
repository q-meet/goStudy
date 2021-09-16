package main

import (
	"fmt"
	"net/http"
	"time"
)
// 声明一个中间件类型
type middleware func(handler http.Handler) http.Handler
// 定义一个路由
type Router struct {
	middlewareChain []middleware
	mux map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		mux:make(map[string]http.Handler),
	}
}
// 添加中间件
func (r *Router) Use (m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}
// 添加路由
func (r *Router) Add(route string, h http.Handler) {
	var mergedHandler = h
	for i := len(r.middlewareChain) - 1; i>=0; i-- {
		// 倒序 循环构造中间件合集， 类似递归的方式 嵌套组合成一个方法
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}
	// 把最终合成的路由处理方法赋值
	r.mux[route] = mergedHandler
}
// 实现 http.Hander 接口
func (r *Router) ServeHTTP(w http.ResponseWriter, q *http.Request) {
	path := q.URL.Path
	if path == "/hello" {
		r.mux[path].ServeHTTP(w, q)
	} else {
		w.Write([]byte("zzh not find"))
	}
}

// time middle
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		timeStart := time.Now().UnixNano() / 1e6

		next.ServeHTTP(w, r)

		timeEnd := time.Now().UnixNano() / 1e6

		timeElapsed := timeEnd - timeStart

		fmt.Println("处理时间:", timeElapsed)
	})
}

func pathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before pathMiddleware")
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Println("after pathMiddleware")
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is hello function")
	_, err := w.Write([]byte("hello zzh"))
	if err != nil {
		fmt.Println(err)
	}
}

func main(){
	r := NewRouter()
	r.Use(timeMiddleware)
	r.Use(pathMiddleware)
	r.Add("/hello", http.HandlerFunc(hello))
	//http.HandleFunc("/hello", hello)
	http.ListenAndServe(":12345", r)
}