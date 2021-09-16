# Job 1
```text
封装一个数据结构 MyMap，
实现并发安全的 Load，Store，Delete，LoadAndDelete，LoadOrStore 几个 API(禁止使用 sync.Map)
，不用考虑性能，文件在 https://golearn.coding.net/p/gonggongbanji/files/all/DF18 中的 mymap.go
```
### 目标文件 mymap.go

# Job 2
```text
编写 benchmark
比较 MyMap 与 sync.Map 的同名函数性能差异(LoadAndDelete 可以不用比较，该函数在 Go 1.15 引入，我们的作业环境是 1.14.12)，
输出相应的性能报告(注意，你应该使用 RunParallel)，将性能比较结果输出为 markdown 文件
```
### 目标文件 mymap_test.go

###### 执行命令
```shell
go test -bench .  >testResult.md
```

###### 查看代码覆盖率 html方式打开
```shell
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

# Job 3
```text
使用 channel 实现一个 trylock，模板 trylock.go
```


# Job 4
```text
修复 https://golearn.coding.net/p/gonggongbanji/files/all/DF18 中的 deadlock.go 的死锁
```

# Job 5
```text
在 context 学习过程中，我们知道，在子节点中 WithValue 的数据，父节点是查不到的。
请实现一个 MyContext，其 WithValue 方法赋值的 k，v 在父节点中也可以查得到。模板：mycontext.go。
```


# Job 6
```text
实现，或封装社区的 timewheel，使用 benchmark 或 pprof 比较大量 timer 存活时，内置 timer 和 timewheel 的性能差异。不要求提交代码，只需要性能对比报告 markdown。
```

