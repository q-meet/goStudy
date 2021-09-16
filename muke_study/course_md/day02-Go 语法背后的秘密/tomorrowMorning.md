阅读完整的 课后pdf文件

完成 Dockerfile作业

```作业
作业
• 必做：channel 一⻚中，找出所有红线报错在 runtime 中的位置，使用 Go 1.14.12  版本
• 选做1：修复https://golearn.coding.net/p/gonggongbanji/files/all/DF12中的 test case
• 选做2：扩展其功能，增加更多的操作符号支持(随意发挥)
• 难度++选做：分析 Go 语言新版本的函数调用规约
• 必做作业，将分析过程总结为 markdown 文档，•提交到 coding 的班级仓库下，带上自己的慕课 ID
```

编译与反编译工具-编译
go tool compile -S ./hello.go | grep “hello.go:5”


dlv 反汇编
[disass](https://github.com/go-delve/delve/tree/master/Documentation/cli)

编译与反编译工具-反编译go tool objdump 寻找 make 的实现

函数调用规约
栈本质上是内存
内置在系统里面核心的寄存器 访问速度快

[dlv](https://zhuanlan.zhihu.com/p/373559087)  完成

[GoCN 上的 dlv 的新译⽂](https://gocn.vip/topics/12090)  

[编译反编译工具调试](https://studygolang.com/articles/18124)  

[dlv](https://zhuanlan.zhihu.com/p/373559087)  完成

[JetBrains GoLand 2021.1 新特性介绍](https://www.bilibili.com/video/BV1Vo4y117q6)

[Go 面试官问我如何实现面向对象？](https://mp.weixin.qq.com/s/2x4Sajv7HkAjWFPe4oD96g)

[dlv 是怎么用的](https://github.com/go-delve/delve)

[dlv 演示](https://github.com/go-delve/delve/blob/master/Documentation/installation/README.md)  

[编译反编译工具调试](https://studygolang.com/articles/18124) √

[曹春晖：谈一谈 Go 和 Syscall](https://blog.csdn.net/cmqsbon24073/article/details/100414185?utm_medium=distribute.pc_relevant_t0.none-task-blog-2~default~BlogCommendFromMachineLearnPai2~default-1.control&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2~default~BlogCommendFromMachineLearnPai2~default-1.control)


[main.main之前的准备](https://www.cntofu.com/book/3/zh/04.2.md) 

词法分析等文件库了解   



[16-go调度器触发-syscall](https://blog.csdn.net/svasticah/article/details/95252646)
[go语法书入门](https://github.com/chai2010/go-ast-book) 


[goroutine 调度时机有哪些](https://qcrao91.gitbook.io/go/goroutine-tiao-du-qi/goroutine-tiao-du-shi-ji-you-na-xie)

[Go 面试官问我如何实现面向对象？](https://mp.weixin.qq.com/s/2x4Sajv7HkAjWFPe4oD96g)
[dlv 是怎么用的](https://github.com/go-delve/delve) 



测试覆盖率 注： -coverprofile 标志自动设置 -cover 来启用覆盖率分析。
go test -covermode=count -coverprofile fib.out

查看的更有趣的方式是获取 覆盖率信息注释的源代码 的HTML展示。 该显示由 -html 标志调用：
go tool cover -html=fib.out -n flb.html

go test 命令接受 -covermode 标志将覆盖模式设置为三种设置之一：

set: 每个语句是否执行？
count: 每个语句执行了几次？
atomic: 类似于 count, 但表示的是并行程序中的精确计数


下面来试试一个标准包， fmt 格式化包语句执行的计数。 进行测试并写出 coverage profile ，以便能够很好地进行信息的呈现。
go test -covermode=count -coverprofile=../src/cover/count.out fmt

这比以前的例子好的测试覆盖率。 （覆盖率不受覆盖模式的影响）可以显示函数细节：
go tool cover -func=../src/cover/count.out

HTML输出产生了巨大的回报：
go tool cover -html=../src/cover/count.out


尽量的运用cpu核心资源 然所有cpu都在执行计算

Go 的词法分析和语法/语义分析过程：https://dev.to/nakabonne/digging-deeper-into-the-analysis-of-go-code-31af

编译器各阶段的简单介绍：
https://www.tutorialspoint.com/compiler_design/compiler_design_phases_of_compiler.htm

Linkers and loaders， 只看内部对 linker 的职责描述就行，不用看原理
https://golearn.coding.net/p/gonggongbanji/files/all/DF9

SSA 的简单介绍(*只做了解)：
https://mp.weixin.qq.com/s/UhxFOQBpW8EUVpFvqH2tMg

老外的写的如何定制 Go 编译器，里面对 Go 的编译过程介绍更详细，SSA 也说明得很好(*只做了解)：
https://eli.thegreenplace.net/2019/go-compiler-internals-adding-a-new-statement-to-go-part-2/

如何阅读 go 的 SSA(*难，只做了解)：
https://sitano.github.io/2018/03/18/howto-read-gossa/


CMU 的编译器课，讲 SSA(*难，只做了解) https://www.cs.cmu.edu/~fp/courses/15411-f08/lectures/09-ssa.pdf

对逆向感兴趣的话(扩展内容，与本课程无关)：https://malwareunicorn.org/#/workshops

Vitess 的 SQL Parser：https://github.com/xwb1989/sqlparserPing

CAP 的 TiDB 的 SQL Parser：https://github.com//pingcap/parser

GoCN 上的 dlv 的新译文https://gocn.vip/topics/12090

C语言调用规约
https://github.com/cch123/llp-trans/blob/master/part3/translation-details/function-calling-sequence/calling-convention.md

Go 语言新版调用规约
https://go.googlesource.com/proposal/+/refs/changes/78/248178/1/design/40724-register-calling.md
