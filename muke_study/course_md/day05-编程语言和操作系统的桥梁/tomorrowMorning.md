
```作业
1-10 第三课作业
必做：
封装一个数据结构 MyMap，实现并发安全的 Load，Store，Delete，LoadAndDelete，LoadOrStore 几个 API(禁止使用 sync.Map)，不用考虑性能，文件在https://golearn.coding.net/p/gonggongbanji/files/all/DF18 中的 mymap.go

选做：编写 benchmark，比较 MyMap 与 sync.Map 的同名函数性能差异(LoadAndDelete 可以不用比较，该函数在 Go 1.15 引入，我们的作业环境是 1.14.12)，输出相应的性能报告(注意，你应该使用 RunParallel)，将性能比较结果输出为 markdown 文件
选做：使用 channel 实现一个 trylock，模板 trylock.go
选做：修复 https://golearn.coding.net/p/gonggongbanji/files/all/DF18 中的 deadlock.go 的死锁
选做：在 context 学习过程中，我们知道，在子节点中 WithValue 的数据，父节点是查不到的。请实现一个 MyContext，其 WithValue 方法赋值的 k，v 在父节点中也可以查得到。模板：mycontext.go。
选做难度++：实现，或封装社区的 timewheel，使用 benchmark 或 pprof 比较大量 timer 存活时，内置 timer 和 timewheel 的性能差异。不要求提交代码，只需要性能对比报告 markdown。
提交方式：
将必做和选做的相关代码和 markdown 文件
提交到 coding 班级仓库中，peer review
对并发编程不太熟悉的同学，先阅读：
https://github.com/gopl-zh/gopl-zh.github.com
的第 8，第 9 章，选做题量力而行
```
