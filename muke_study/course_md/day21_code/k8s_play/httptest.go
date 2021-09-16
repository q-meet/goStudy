package main

import (
	"io"
	_ "net/http/pprof"

	"net/http"
)

func sayhello(wr http.ResponseWriter, r *http.Request) {
	io.WriteString(wr, "hello3")
}

/*
作业
必做：
	使⽤ docker hub 发布⾃⼰的 hello 服务
	使⽤ minikube 将代码打包成为 deployment，并发布为 service
	在本机访问 minikube 中的服务(提示：minikube tunnel 或 service 命令)
选做：
	在 minikube 中集成服务发现，打通服务间调⽤(提示，默认的服务发现可以通
	过 env 找到其它服务的 host 和 port；愿意选择 consul 来做也可以)
	在 minikube 上配置定时任务
*/
func main() {
	http.HandleFunc("/", sayhello)
	http.HandleFunc("/hello", sayhello)
	http.ListenAndServe(":11114", nil)
}
