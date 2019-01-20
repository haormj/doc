### Protobuf

文档

https://developers.google.com/protocol-buffers/

首先需要安装protobuf的编译器,下载配置环境变量即可

https://github.com/google/protobuf

然后继续安装一个go的插件

https://github.com/golang/protobuf/

```shell
go get -u github.com/golang/protobuf/protoc-gen-go
```

编译proto文件

```shell
protoc --go_out=. *.proto
```

