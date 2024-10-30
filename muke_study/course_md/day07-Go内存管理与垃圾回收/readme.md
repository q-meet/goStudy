# 内存分配与垃圾回收

## 内存分配

1. 基本概念  
  从问题出发
   - 内存从哪里来
   - 内存到哪里去

   - 标记对象从哪里来
   - 标记对象到哪里去

   - 垃圾从哪里来
   - 垃圾到哪里去

栈分配是非常轻量的工作 使用完函数栈帧(sp下移)就会自动销毁

### 如何找到所有逃逸分析的可能性  

- 高难度模式 看源代码文件
  cmd/compile/internal/gc/escape.go

- 低难度模式 看官方demo
<https://github.com/goland/go/tree/master/test>  
  
有内存分配的语言:自动allocator go自动分配  
有垃圾回收的语言:自动collector 自动回收  
自动内存回收技术也叫垃圾回收技术

- 内存管理的中的三个角色
  - Mutator:fancy(花哨的) word for application, 其实就是自己写的应用程序，它会不断地修改对象的引用关系，即对象图
  - Allocator:内存分配器，负责管理从操作系统中分配出的内存空间，malloc 其实底层就有一个内存分配器的实现(glibc 中)，tcmalloc 是malloc 多线程改进版。Go 中的实现类似 tcmalloc
  - Collector：垃圾收集器，负责清理死对象，释放内存空间

内存管理抽象
每个操作系统都有对象的实现
mem_linux.go
mem_windows.go
相关的抽象描述在 runtime/malloc.go

2. Allocator基础
   Bump 分配器
   Free List Allocator 分配器  
   动画:  
   <https://www.figma.com/proto/tSl3CoSWKitJtvIhqLd8Ek/memory-management-and-and-garbage-collection?page-id=175%3A118&node-id=175%3A119&viewport=-7626%2C499%2C0.4998303949832916&scaling=contain>
   Free List Allocator 分配器详解 空闲链表 大致有以下几种类型
   - First-Fit
   - Next-Fit
   - Best-Fit
   - Segregated-Fi  不同大小分区（go采用类似的算法实现）
    动画地址:
     <https://www.figma.com/proto/tSl3CoSWKitJtvIhqLd8Ek/memory-management-and-and-garbage-collection?page-id=233%3A21&node-id=233%3A22&viewport=-650%2C182%2C0.059255365282297134&scaling=min-zoom>
3. malloc实现
4. Go语言内存分配  
    动画地址:  
   <https://www.figma.com/proto/tSl3CoSWKitJtvIhqLd8Ek/memory-management-and-and-garbage-collection?page-id=151%3A36&node-id=151%3A38&viewport=241%2C543%2C0.2360718995332718&scaling=contain>

    - 栈内存分配
      分配大小分类：•Tiny : size < 16 bytes && has no pointer(noscan)
      Small ：has pointer(scan) || (size >= 16 bytes && size <= 32 KB)
      Large : size > 32 KB
    - 堆内存分配

## 垃圾回收

1. 垃圾回收基础  
   - 垃圾分类  
       - 语义垃圾--有的被称为内存泄露
       语义垃圾值得是从语法上可达（可以通过局部、全局变量引用得到）的对象，但从语义上来讲他们是垃圾，垃圾回收器对此无能为力  
       动画地址：
     <https://www.figma.com/proto/tSl3CoSWKitJtvIhqLd8Ek/memory-management-and-and-garbage-collection?page-id=185%3A2&node-id=185%3A13&viewport=-45%2C330%2C0.1481591761112213&scaling=contain>  
        - 语法垃圾
        语法垃圾是讲那些从语法上无法到达的对象，这些才是垃圾收集器主要的搜集目标

2. Go语言垃圾回收
     垃圾收集算法可视化:<https://spin.atomicobject.com/2014/09/03/visualizing-garbage-collection-algorithms/>
   - 常⻅垃圾回收算法
     引用计数(Reference Counting)：某个对象的根引用计数变为 0 时，其所有子节点均需被回收。
     标记压缩(Mark-Compact)：将存活对象移动到一起，解决内存碎片问题。
     复制算法(Copying)：将所有正在使用的对象从 From 复制到 To 空间，堆利用率只有一半。
     标记清扫(Mark-Sweep)：解决不了内存碎片问题。需要与能尽量避免内存碎片的分配器使用，如 tcmalloc。<— Go 在这里

垃圾回收入口：gcStart

手工调用 runtime.GC函数(一般测试GC bug的时候使用)  
内存分配的时候 runtime.mallcgc (分配速度超过回收速度 会导致内存一直分配下去最后会导致oom
)  
后台的gc触发的goroutime forcegchelper(垃圾回收没有及时触发 就主动触发 2min 没有触发gc就会触发)
3. Gc标记流程
    - GC 标记流程-三色抽象

    - 三色抽象
        * 黑：已经扫描完毕，子节点扫描完毕。(gcmarkbits = 1，且在队列外。)
        * 灰：已经扫描完毕，子节点未扫描完毕。(gcmarkbits = 1, 在队列内)
        * 白：未扫描，collector 不知道任何相关信息
    - 动画地址
      https://www.figma.com/proto/tSl3CoSWKitJtvIhqLd8Ek/memory-management-and-and-garbage-collection?page-id=0%3A1&node-id=2%3A38&viewport=124%2C371%2C0.11918419599533081&scaling=contain

GC 时要注意的问题：
1.对象在标记过程中不能丢失
2.Mark 阶段 mutator 的指向堆的指针修改需要被记录下来
3.GC Mark 的 CPU 控制要努力做到 25% 以内

强三色不变性
strong tricolor invariant禁止黑色对象指向白色对象

弱三色不变性
weak tricolor invariant黑色对象指向的白色对象，如果有灰色对象到它的可达路径，那也可以
