package main

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var task = map[string]string{
	"bing.html":   "https://www.bing.com",
	"baidu.html":  "https://www.baidu.com",
	"google.html": "https://www.google.com",
}

/*
使用任意并发工具，完成这样一个程序，
并发请求 https://baidu.com 和 https://bing.com，
任意一个站点先获取到了完整的 html，即中止另一个流程，并把结果输出至文件：
{sitename}.html，例如百度先获取到了结果，即输出baidu.html，若 bing 先获取到了结果，就输出 bing.html。
*/
func main() {
	// 根 context
	ctx, cancel := context.WithCancel(context.Background())
	// 流程控制 chan
	CancelTask := make(chan struct{})
	// 循环请求
	for k, v := range task {
		go func(k, v string) {
			resp, respErr := requestHttp(ctx, CancelTask, v)
			if respErr != nil {
				log.Println("发生错误", respErr)
				return
			}
			//把文件 存入
			res, err := writeContext(k, resp)
			if err != nil {
				log.Println("文件存储发生错误", respErr)
				return
			}
			if res {
				log.Println(k, ":文件写入成功")
			}
			CancelTask <- struct{}{}
		}(k, v)
	}
	// 第一个完成
	<-CancelTask
	// 取消其他任务
	cancel()
	//完成存入
	<-CancelTask

	//查看取消流程
	time.Sleep(time.Second * 2)

}

// writeContext 创建文件且写入内容
func writeContext(path string, context []byte) (bool, error) {
	if context == nil {
		return false, errors.New("context be empty")
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return false, err
	}

	// 关闭文件
	defer file.Close()

	//内容入库
	_, err = file.Write(context)
	if err != nil {
		return false, err
	}
	return true, nil
}

// requestHttp 请求文件内容 并返回
func requestHttp(ctx context.Context, CancelTask chan<- struct{}, url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("url be empty")
	}
	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	//关闭连接
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("resp status code err" + string(resp.StatusCode))
	}
	//读取返回内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 任务取消 或 结果返回
	select {
	case CancelTask <- struct{}{}:
		return body, nil
	case <-ctx.Done():
		return nil, errors.New("task cancel")
	}
}
