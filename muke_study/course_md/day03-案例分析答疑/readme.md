# 答疑

###  day1
- 部署好本机的docker环境，使用ppt中的dockerfile build自己的环境

- 使用 readelf 工具，查看编译后的进程入口地址
```shell
//打包为可执行文件
go build main.go
```

```shell
//使用readelf查看 Entry point address
[root@fdd4fd454c8e home]# go build
[root@fdd4fd454c8e home]#
[root@fdd4fd454c8e home]# ls
home  main.go
[root@fdd4fd454c8e home]# readelf -h ./home
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              EXEC (Executable file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x455780
  Start of program headers:          64 (bytes into file)
  Start of section headers:          456 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         7
  Size of section headers:           64 (bytes)
  Number of section headers:         25
  Section header string table index: 3
```

- 使用dlv 执行这个可执行文件 且打一个断点
    * c  // 运行程序直到程序结束或遇到下一个端点
    * n  // 执行源文件的下一行
    * s  // 单步执行。如果遇到函数调用，则进入到被调用的函数中。和next的区别是，当next遇到函数调用时，不进入函数内部，仍留在主函数中。具体例子中会讲解。
    * si  //单步执行cpu指令。执行一条汇编指令
    * so   //单步跳出函数，返回到调用函数的那一行。
    * disass   //查看当前执行反编译命令到哪行了
      
    * b   //设置一个断点
    * bp  //打印出当前所有的断点
    * condition (alias: cond)  Set breakpoint condition. //设置断点条件
    * list (alias: ls | l) ------- Show source code. //查看源代码
    * stack (alias: bt)  Print stack trace. 函数从什么地方跳进来的 查看调用关系
    * frame ------------ Set the current frame, or execute command on a different frame. 跳转bt位置
    
```shell
//dlv exec home
[root@fdd4fd454c8e home]# dlv exec ./home
Type 'help' for list of commands.
(dlv) b *0x455780
Breakpoint 1 set at 0x455780 for _rt0_amd64_linux() /usr/local/go/src/runtime/rt0_linux_amd64.s:8

(dlv)c
(dlv)si
(dlv)si
(dlv)
```



rbp 函数栈底
rsp 函数栈顶