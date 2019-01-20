## redis aof占用磁盘过高

### 实际现象

1. 磁盘占用很高
2. redis-server日志报错
3. 手动执行`BGREWRITEAOF`失败
4. 手动执行`BGSAVE`失败

### 问题定位

#### 版本

redis_version:3.0.6

#### 日志

```
1599:M 21 Nov 21:35:41.630 * Starting automatic rewriting of AOF on 96% growth
1599:M 21 Nov 21:35:41.630 # Can't rewrite append only file in background: fork: Cannot allocate memory
1599:M 21 Nov 21:38:21.416 * Starting automatic rewriting of AOF on 103% growth
1599:M 21 Nov 21:38:21.416 # Error opening /setting AOF rewrite IPC pipes: Numerical result out of range
```

#### 结论

当前版本的redis在写入aof的时候,当fork失败,会导致pipes泄露

### 解决方案

#### 思路

关闭aof,使用rdb

#### 步骤

##### 执行BGSAVE

```
Can't save in background: fork: Cannot allocate memory
```

##### 修改系统参数

```shell
# 临时
sysctl -w vm.overcommit_memory=1
# 永久通过编辑 /etc/sysctl.conf
# 然后让其生效
sysctl -p
```

##### 设置redis参数

```shell
redis-cli -h host -p port config set appendonly no
redis-cli -h host -p port CONFIG SET SAVE "900 1 300 10 60 10000"
# 在配置文件中永久配置,防止redis重启之后失效
```

### 参考

1. [Trouble Shooting —— Redis AOF rewrite错误导致Redis被Block住](https://ningyu1.github.io/site/post/11-redis-aof-pit/)
2. [Error opening /setting AOF rewrite IPC pipes: Numerical result out of range](https://github.com/antirez/redis/issues/2857)
3. [Redis 源码 bug 深入定位过程分享](https://cloud.tencent.com/developer/article/1005716)
4. [Background saving fails with a fork() error under Linux even if I have a lot of free RAM!](https://redis.io/topics/faq)

