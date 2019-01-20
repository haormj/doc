## redis-port

### 简介

[redis-port](https://github.com/CodisLabs/redis-port)可以全量+增量将数据同步到另一个redis,做到数据的平滑迁移.将自己伪装为副本,从主同步数据,依赖`sync`/`psync`命令.

### 下载

直接到releases页面,下载指定版本即可,这里使用的是[v2.0-beta](https://github.com/CodisLabs/redis-port/releases/tag/v2.0-beta)

### 使用

#### 同步

##### 运行

```shell
./redis-sync \
  --master=host:6379 \
  --target=host:6379 \
  --db=* \
  --ncpu=8 \
  --tmpfile-size=10gb \
  --tmpfile=./sockfile.tmp
```

##### 日志

```
2018/06/21 11:52:42 sync.go:76: [INFO] sync: master = "host:6379", target = "host:6379"
2018/06/21 11:52:43 sync.go:103: [INFO] +
2018/06/21 11:52:44 sync.go:105: [INFO] -
2018/06/21 11:52:44 sync.go:103: [INFO] +
2018/06/21 11:52:45 sync.go:109: [INFO] sync: runid = "990718ada62b4efe861c7cf08fcd65e1b5655418", offset = 13409028
2018/06/21 11:52:45 sync.go:110: [INFO] sync: rdb file = 75741891 (72.23mb)
2018/06/21 11:52:45 sync.go:208: [INFO] sync: (r/f,s/f,s) = (read,rdb.forward,rdb.skip/rdb.forward,rdb.skip)
2018/06/21 11:52:46 sync.go:250: [INFO] sync: rdb = 75741891 - [100.00%]   (r/f,s/f,s)=(3895664/36459,0/0,0)     ~  (3.72mb/-,-/-,-)  ~  speed=(3.72mb/36459,0/0,0)
```

##### 解释

```
r master.rbytes.Int64()
f master.rdb.forward.Int64()
s master.rdb.skip.Int64()
f master.aof.forward.Int64()
s master.aof.skip.Int64()
```

