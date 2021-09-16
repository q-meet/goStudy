## 编译与反编译工具

### 编译与反编译工具-编译
该命令会生成 .o 目标文件，并把目标的汇编内容输出
- go tool compile -S 1.go | grep "1.go:8"

- go tool compile -S str_to_byte.go | grep "str_to_byte.go:5"


### 编译与反编译工具-反编译
- go tool objdump 寻找 make 的实现https://golang.org/ref/spec，官方 spec


```shell
go build make.go && go tool objdump ./make | grep -E "make.go:6|make.go:10|make.go:14"
```

```shell
go build -gcflags="-N -l" new.go && go tool objdump new | grep -E "new.go:6|new.go:7|new.go:8|new.go:9"
```

- 优化之后的

```shell
go tool compile -N -S  new.go | grep -E "new.go:6|new.go:7|new.go:8|new.go:9"
```

这里使用 go tool compile -S 也是可以的。

观察 new(**输出内容难以读懂，不推荐**)


