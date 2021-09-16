阅读完整的 课后pdf文件

完成 Dockerfile作业

```作业
作业：
* 部署好本机的 docker 环境，使用 ppt 中的 dockerfile build 自己的环境
* 使用 readelf 工具，查看编译后的进程入口地址 || readelf -h ./go
* 在 dlv 调试工具中，使用断点功能找到代码位置[dlv基础使用](https://github.com/chai2010/advanced-go-programming-book/blob/master/ch3-asm/ch3-09-debug.md)
* 使用断点调试功能，查看 Go 的 runtime 的下列函数执行流程，使用 IDE 查看函数的调用方：
    * 必做：runqput，runqget，globrunqput，globrunqget
    * 选做：schedule，findrunnable，sysmon
* 难度++课外作业：跟踪进程启动流程中的关键函数，rt0_go，需要汇编知识，可以暂时不做，只给有兴趣的同学

docker build -t test .




```

61魔术变量(完成)


sudog

### go runtime
    可以认为 runtime 是为了实现额外的功能，而在程序运行时自动加载/运行的一些模块。



清晰GPM的一套  
[码农桃花源文章](https://qcrao91.gitbook.io/go/goroutine-tiao-du-qi) 进行中

goroutine的调度  
[Go夜读](https://www.bilibili.com/video/BV1pb411v7nu?t=3177) ok


coding CI CD 尽量完成流程部署

[100 如何高效的阅读 Go 代码？ go夜读](https://www.bilibili.com/video/BV1XD4y1U7Pf)进行中

[86 Go unsafe pointer 使用规则详解【Go 夜读】](https://www.bilibili.com/video/BV15V411d7WS) 完成


理解可执行文件ELF

汇编游戏
人力资源机器

### GoCN 每日新闻 (2021-05-17)

1. 如何管理多版本的go   https://lakefs.io/managing-multiple-go-versions-with-go/
2. 一个SQL数据库只使用了2000行golang代码并且没有任何第三方依赖 https://github.com/auxten/go-sqldb
3. 使用go开发一个监控剪贴板的服务  https://www.reddit.com/r/golang/comments/ncsnj2/monitoring_clipboard_as_service_with_go/
4.  Gorm 复杂关系举例：  https://github.com/harranali/gorm-relationships-examples
5. 写了500，000行go代码之后  https://blog.khanacademy.org/half-a-million-lines-of-go
    * GopherChina 全部日程出炉 https://gc.gocn.vip

- 编辑: 阿章
- 订阅新闻: http://tinyletter.com/gocn
- 招聘专区: https://gocn.vip/jobs
- GoCN归档: https://gocn.vip/topics/12077