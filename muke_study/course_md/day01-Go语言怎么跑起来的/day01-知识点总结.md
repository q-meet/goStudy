# 知识点总结

## 可执行文件 ELF:
* 使用 go build -x 观察编译和链接过程
* 通过 readelf -H 中的 entry 找到程序入口
* 在 dlv 调试器中 b *entry_addr 找到代码位置

## 启动流程：
* 处理参数 -> 初始化内部数据结构 -> 主线程 -> 启动调度循环

## Runtime 构成：
* Scheduler、Netpoll、内存管理、垃圾回收


##GMP:
	M(线程),任务消费者; 
    G(goroutine),计算任务; 
    P(token(上下文环境)),可以使用的CPU的token

## 队列:
P的本地runnext字段->P的local run queue-> global run queue, 多级队列减少锁竞争

## 调度循环：
线程M在持有P的情况下不断消费运行队列中的G的过程

##处理阻塞
* 可以接管的阻塞: channel收发, 网络链接/读写, 加锁, select
* 不可接管的阻塞: syscall, cgo, 长时间运行需要剥离P执行

## sysmon
* 一个后台高优先级循环, 执行时不需要绑定任何P
* 负责
    * 检查是否已经没有活动线程, 如果是则崩溃
    * 轮询 netpoll
    * 剥离在syscall上阻塞的M的P
    * 发信号,抢占已经执行时间过长的G

###
    调度的机制用一句话描述:
    runtime准备好G,P,M, 然后M绑定P, M从各种队列中获取G, 切换到G的执行栈上并执行G上的任务函数, 调用goexit 做清理工作并回到M, 如此反复


### 展开说明 基本概念
##### M(Machine)
* M代表着真正的执行计算资源，可以认为他就是os thread(系统线程)
* M是真正调度系统的执行者，每个M就像一个勤劳的工作者，总是从各种队列中找到可运行的G, 而且这样M的可以同时存在多个
* M在绑定有效的P之后，进度调度循环，而且M并不保留G的状态，这是G可以跨M调度的基础

##### P(processor)
* P表示逻辑 processor，是线程M的执行的上下文
* P的最大作用是其拥有的各种G对象队列、链表、cache和状态

##### G(goroutine) 
* 调度系统的最基本单元 goroutine ， 存储了goroutine的执行stack(堆)信息、goroutine状态以及goroutine的任务函数等。
* 在G眼中只有P，P就是运行G的"CPU"。
* 类似于两级线程



## 队列:
	goroutine处理有三个队列(多级队列减少锁竞争)

	P的本地runnext字段(只能存储一个)

	P的local run queue(数组固定大小256)

	global run queue(切片 无限大小)

## 调度循环：
	线程M在持有P的情况下不断消费运行队列中的G的过程

## 处理阻塞:
	1. 可以接管的阻塞：channel收发, 加锁, 网络连接读/写, select
	2. 不可接管的阻塞: syscall, cgo, 长时间运行需要剥离P执行


注意，要尽量讲得浅显易懂
要注意引申，就是在别的什么地方有类似的设计啊，或者说可能的解决思路还有啥的
站在设计者的角度，来给我们分享源码


sysmon处理(cgo,syscall)阻塞协程  不使用p直接执行 高优先级

netpoll 会再sysmon里面执行一次


b procresize

continue

bt

断点:
b runqput
c
参数
locals
args
l
n 下一步



在什么位置输出错误