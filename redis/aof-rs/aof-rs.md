## aof-rs

### 介绍

aof文件过滤工具,可以过滤cmd/key/db,地址 https://github.com/bosondata/aof-rs

### 使用

```shell
# -k 后面可以为正则表达式
./aof -k world*  appendonly.aof.bak 

# 匹配指定前缀的key
./aof -k '^(hello|world).*$' appendonly.aof.bak 
```

