## Redis

临时开启aof

```shell
redis-cli -h host -p port config set appendonly yes
```

临时关闭aof

```shell
redis-cli -h host -p port config set appendonly yes
```

导入aof文件

```
redis-cli -h host -p port -a password --pipe < appendonly.aof
```

查看key的数量

```shell
127.0.0.1:6379> DBSIZE
(integer) 3246711
```

动态设备rdb保存

```shell
redis-cli -h host -p port CONFIG SET SAVE "900 1 300 10 60 10000"
```

























