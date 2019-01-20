### cgroup

#### cpu

```
cgroup.clone_children
cgroup.event_control
cgroup.procs
cpuacct.stat
cpuacct.usage
cpuacct.usage_percpu
# 时间周期长度
cpu.cfs_period_us
# 在当前cgroup在设置的周期长度内所能使用的cpu时间数
cpu.cfs_quota_us
cpu.rt_period_us
cpu.rt_runtime_us
# cpu的相对值
cpu.shares
# cpu统计值
cpu.stat
notify_on_release
tasks
```

#### memory

```
cgroup.clone_children 
cgroup.event_control
# 进程id
cgroup.procs
# 线程id
tasks
# 内存使用量达到上限值的次数
memory.failcnt 
# 触发系统尽可能回收当前cgroup可以回收的内存
memory.force_empty
# 限制内存的大小
memory.limit_in_bytes 
# 历史内存最大使用量
memory.max_usage_in_bytes 
memory.move_charge_at_immigrate
memory.numa_stat
memory.oom_control
memory.pressure_level
# 当前限制的内存软额度,当系统内存吃紧时,会回收超过soft limit的内存
memory.soft_limit_in_bytes
# 当前内存使用情况
memory.stat 
memory.swappiness
# 当前已使用内存
memory.usage_in_bytes
# 是否将子cgroup的使用情况统计到当前cgroup中
memory.use_hierarchy
notify_on_release
# memsw
memory.memsw.failcnt 
memory.memsw.limit_in_bytes
memory.memsw.max_usage_in_bytes
memory.memsw.usage_in_bytes
# kmem
memory.kmem.failcnt
memory.kmem.limit_in_bytes
memory.kmem.max_usage_in_bytes 
memory.kmem.slabinfo
memory.kmem.tcp.failcnt
memory.kmem.tcp.limit_in_bytes
memory.kmem.tcp.max_usage_in_bytes
memory.kmem.tcp.usage_in_bytes
memory.kmem.usage_in_bytes
```

#### 参考

1. [ Linux Cgroup系列（04）：限制cgroup的内存使用（subsystem之memory）](https://segmentfault.com/a/1190000008125359)
2. [Linux Cgroup系列（05）：限制cgroup的CPU使用（subsystem之cpu）](https://segmentfault.com/a/1190000008323952)