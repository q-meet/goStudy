# 并发编程与最佳实践

## 并发内置数据结构

### sync

sync.Once

    sync.Once 只有一个方法，Do()
    但o.Do需要保证：
    1. 初始化方法必须且只能被调用一次
    2. Do返回后，初始化一定已经执行完成
sync.Pool

    主要在两种场景使用:
    1. 进程中的inuse_objects（应用）数过多，gc mark消耗大量CPU
    2. 进程中的inuse_objects数过多，进程RSS占用过高
    请求生命周期开始时，pool.Get, 请求结束时,pool.Put。
    在fasthttp中有大量应用.
    https://github.com/valyala/fasthttp/blob/b433ecfcbda586cd6afb80f41ae45082959dfa91/server.go#L402

sync.Mutex
sync.RWMutex
sync.Map
sync.Waitgroup

## 并发编程模式举例

## 常见的并发bug

死锁
map并发读写

Channel 关闭 panic
Channel closing principle

## 内存模型
