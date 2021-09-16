package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
)

type product struct {
	Index    uint64  `json:"index"`           // 用户
	Title    string  `json:"title"`           // 用户
	SubTitle string  `json:"sub_title"`       // 微博内容
	Pic      string  `json:"pic"`             // 转发数
	Price    float64 `json:"price,omitempty"` // 图片
}

// 索引mapping定义，这里仿微博消息结构定义
const mapping = `
{
  "mappings": {
    "properties": {
		"index": {
		  "type": "long"
		},
      "title": {
        "type": "text"
      },
      "sub_title": {
        "type": "text"
      },
      "price": {
        "type": "text"
      },
      "pic": {
        "type": "text"
      }
    }
  }
}`

var client *elastic.Client
var ctx context.Context
var err error

func init() {

	// 创建client
	client, err = elastic.NewClient(
		// elasticsearch 服务地址，多个服务地址使用逗号分隔
		elastic.SetURL("http://127.0.0.1:9200"),
		// 基于http base auth验证机制的账号和密码
		elastic.SetBasicAuth("", ""),
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置请求失败最大重试次数
		elastic.SetMaxRetries(5),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}

	// 执行ES请求需要提供一个上下文对象
	ctx = context.Background()

	// 首先检测下索引是否存在
	exists, err := client.IndexExists("canal_product4").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// 索引不存在，则创建一个
		_, err := client.CreateIndex("canal_product4").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
	}
}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}

	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", res.TookInMillis, res.TotalHits())
	// 查询结果不为空，则遍历结果
	var typ product
	// 通过Each方法，将es结果的json结构转换成struct对象
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		// 转换成Article对象
		t := item.(product)
		fmt.Printf("%#v\n", t)
	}
}

/*
使用https://github.com/olivere/elastic完成 MySQL 的数据导入到 elasticsearch 中的需求。

必做：
MySQL 有一张表，你要把 MySQL 的数据插入到 elasticsearch 里，并且能够搜索得到，提供 API。

选做：MySQL 有一张表，并且会实时更新。你需要把更新的数据在 elasticsearch 中也能查询到。MySQL -> Binlog  -> kafka -> consumer -> elasticsearchEs —> query -> keyword -> 搜索结果

作业目的：了解企业中的异构数据同步技术栈，日常工作碰到类似场景可以给出解决方案不用提交
*/
func main() {
	//短语搜索 搜索about字段中有 rock climbing
	//matchPhraseQuery := elastic.NewMatchPhraseQuery("title", "小米8")
	matchPhraseQuery := elastic.NewWildcardQuery("title", "oppo*")
	res, err := client.Search("canal_product4").Query(matchPhraseQuery).Do(ctx)
	printEmployee(res, err)

	//取所有
	res, err = client.Search("canal_product4").Do(context.Background())
	printEmployee(res, err)

	// 创建创建一条微博
	msg1 := product{Index: 2, Title: "oppo1 reno", SubTitle: "打酱油我的天", Price: 522, Pic: "522"}

	// // 使用client创建一个新的文档
	_, err = client.Index().
		Index("canal_product4"). // 设置索引名称
		Id("2").                 // 设置文档id
		BodyJson(msg1).          // 指定前面声明的微博内容
		Do(ctx)                  // 执行请求，需要传入一个上下文对象

	if err != nil {
		// Handle error
		panic(err)
	}
}
