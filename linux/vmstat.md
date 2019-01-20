## vmstat

### 使用

#### 每间隔一秒输出信息

```
vmstat 1
procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
 r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
 0  0 4451508 4271164  66452 1529124    3    7    23    24    5    7  2  1 97  0  0
 0  0 4451508 4269288  66452 1529940    0    0     0     0  946 1612  1  1 98  0  0
```

### 解释

#### procs

r: 正在运行和等待运行的进程数量

b: 在不可中断状态睡眠的进程数量

#### memory(单位:Byte)

swpd: 虚拟内存使用大小

free: 空闲内存大小

buff: 被作为buffer的内存大小

cache: 被作为cache的内存大小

#### swap

si: 从磁盘交换的内存量

so: 交换到磁盘的内存量

#### io

bi: 从块设备接收的块

bo:发送到块设备的块

#### system

in: 每秒的中断数量

cs: 每秒上下文切换的数量

#### cpu(共100%)

us: 用户占用cpu的百分比

sy: 系统占用cpu的百分比

id: 空闲cpu百分比

wa: 等待IO占用cpu的百分比

st: cpu被偷走的百分比

### 参考

1. `man vmstat`

