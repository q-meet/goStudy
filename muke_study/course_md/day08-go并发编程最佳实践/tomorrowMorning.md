## 明日计划

```作业
必做：
使用任意并发工具，完成这样一个程序，并发请求https://baidu.com和https://bing.com，任意一个站点先获取到了完整的 html，即中止另一个流程，并把结果输出至文件：{sitename}.html，例如百度先获取到了结果，即输出baidu.html，若 bing 先获取到了结果，就输出 bing.html。

选做，难度+：安装 herdtools(https://github.com/herd/herdtools7)，在本机上使用 litmus7 执行下列脚本：X86 OOO{ x=0; y=0; } P0          | P1          ; MOV [x],$1  | MOV [y],$1  ; MOV EAX,[y] | MOV EAX,[x] ;locations [x;y;]exists (0:EAX=0 /\ 1:EAX=0)并阅读明白输出结果，无需提交。

选做，难度+++：
回答：为什么使用 atomic.cas 可以实现一个互斥锁，为什么临界区内的内存读写操作不会被重排到 cas 操作之外？
```
