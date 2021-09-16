package main

import (
	"fmt"
	"io"
	_ "net/http/pprof"

	"net/http"
)

var m = map[[12]byte]int{}

//var m = map[string]int{}

func init() {
	for i := 0; i < 1000000; i++ {
		var key [12]byte
		copy(key[:], fmt.Sprint(i))
		m[key] = i
		// m[fmt.Sprint(i)] = i
	}
}

func sayhello(wr http.ResponseWriter, r *http.Request) {
	io.WriteString(wr, "hello")
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":11111", nil)
	if err != nil {
		fmt.Println(err)
	}
}