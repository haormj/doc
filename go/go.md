### Go

go一些好用包的集合

https://github.com/avelino/awesome-go

当写一些通用的包的时候,如何处理日志

https://www.0value.com/

```go
// LogFunc is a function that logs the provided message with optional
// fmt.Sprintf-style arguments. By default, logs to the default log.Logger.
var LogFunc func(string, ...interface{}) = log.Printf
```



golang 类型别名

http://colobu.com/2017/06/26/learn-go-type-aliases/

gogoprotobuf

https://github.com/gogo/protobuf

golang自定义导入包路径

#### go tool pprof

查看cpu

```shell
# cpu
go tool pprof http://xx.xx.xx.xx:xx/debug/pprof/profile
# top占用cpu从高到低的函数
top
# web 可以在浏览器中查看具体内容
web
# svg 可以将信息保存在svg图像中
svg
```

查看内存

```shell
# 内存
go tool pprof http://xx.xx.xx.xx:xx/debug/pprof/heap?debug=1
# 进入之后操作同上
```

