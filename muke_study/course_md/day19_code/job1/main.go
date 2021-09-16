package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"reflect"
	"time"
)

type Weibo struct {
	User     string                `json:"user"`               // 用户
	Message  string                `json:"message"`            // 微博内容
	Retweets int                   `json:"retweets"`           // 转发数
	Image    string                `json:"image,omitempty"`    // 图片
	Created  time.Time             `json:"created,omitempty"`  // 创建时间
	Tags     []string              `json:"tags,omitempty"`     // 标签
	Location string                `json:"location,omitempty"` //位置
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

// 索引mapping定义，这里仿微博消息结构定义
const mapping = `
{
  "mappings": {
    "properties": {
      "user": {
        "type": "keyword"
      },
      "message": {
        "type": "text"
      },
      "image": {
        "type": "keyword"
      },
      "created": {
        "type": "date"
      },
      "tags": {
        "type": "keyword"
      },
      "location": {
        "type": "geo_point"
      },
      "suggest_field": {
        "type": "completion"
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
		elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9200"),
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
}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}

	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", res.TookInMillis, res.TotalHits())
	// 查询结果不为空，则遍历结果
	var typ Weibo
	// 通过Each方法，将es结果的json结构转换成struct对象
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		// 转换成Article对象
		t := item.(Weibo)
		fmt.Printf("%#v\n", t)
	}
}

func main() {

	/*	// 创建创建一条微博
		msg1 := Weibo{User: "olivere22", Message: "打酱油的二天", Retweets: 0}

		// 使用client创建一个新的文档
		_, err := client.Index().
			Index("weibo"). // 设置索引名称
			Id("3"). // 设置文档id
			BodyJson(msg1). // 指定前面声明的微博内容
			Do(ctx) // 执行请求，需要传入一个上下文对象
	*/
	//短语搜索 搜索about字段中有 rock climbing
	//matchPhraseQuery := elastic.NewMatchPhraseQuery("user", "olivere22")
	matchPhraseQuery := elastic.NewWildcardQuery("user","olivere*")
	res, err := client.Search("weibo").Query(matchPhraseQuery).Do(ctx)
	printEmployee(res, err)

	/*
		// 创建terms查询条件
		termsQuery := elastic.NewTermsQuery("user", "olivere", "olivere22")

		searchResult, err := client.Search().
			Index("weibo").   // 设置索引名
			Query(termsQuery).   // 设置查询条件
			Sort("created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
			From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
			Size(10).   // 设置分页参数 - 每页大小
			Do(ctx)             // 执行请求

		printEmployee(searchResult, err)*/

}

func use() {

	// 首先检测下weibo索引是否存在
	exists, err := client.IndexExists("weibo").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// weibo索引不存在，则创建一个
		_, err := client.CreateIndex("weibo").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
	}
	// 创建创建一条微博
	msg1 := Weibo{User: "olivere", Message: "打酱油的二天", Retweets: 0}

	// 使用client创建一个新的文档
	put1, err := client.Index().
		Index("weibo"). // 设置索引名称
		Id("1"). // 设置文档id
		BodyJson(msg1). // 指定前面声明的微博内容
		Do(ctx) // 执行请求，需要传入一个上下文对象
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)

	// 根据id查询文档
	get1, err := client.Get().
		Index("weibo"). // 指定索引名
		Id("1"). // 设置文档id
		Do(ctx) // 执行请求
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("文档id=%s 版本号=%d 索引名=%s\n", get1.Id, get1.Version, get1.Index)
	}

	// 手动将文档内容转换成go struct对象
	msg2 := Weibo{}
	// 提取文档内容，原始类型是json数据
	data, _ := get1.Source.MarshalJSON()
	// 将json转成struct结果
	json.Unmarshal(data, &msg2)
	// 打印结果
	fmt.Println(msg2.Message)

	//根据文档id更新内容
	_, err = client.Update().
		Index("weibo"). // 设置索引名
		Id("1"). // 文档id
		Doc(map[string]interface{}{"retweets": 0}). // 更新retweets=0，支持传入键值结构
		Do(ctx) // 执行ES查询
	if err != nil {
		// Handle error
		panic(err)
	}

	// 根据id删除一条数据
	/*	_, err = client.Delete().
			Index("weibo").
			Id("1").
			Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}*/

}
