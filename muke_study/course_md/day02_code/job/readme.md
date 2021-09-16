## 作业1

channel 一⻚中，找出所有红线报错在 runtime 中的位置，使用 Go 1.14.12  版本  
编译与反编译工具-编译

使用命令查看行数
```shell
cat -n send_to_nil.go
```
设置报错行位置 查看调用函数
```shell
go tool compile -S send_to_nil.go | grep "send_to_nil.go:6"
....
0x0025 00037 (send_to_nil.go:5) CALL    runtime.closechan(SB)  
```
关键行  
得知调用 runtime.closechan 查看函数 得知报错位置


## 作业2
- 选做1：修复https://golearn.coding.net/p/gonggongbanji/files/all/DF12中的 test case
- 选做2：扩展其功能，增加更多的操作符号支持(随意发挥)

文件位置 ./rule_match_test.go