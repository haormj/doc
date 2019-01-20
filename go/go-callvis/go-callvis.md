## go-callvis

github地址

https://github.com/TrueFurby/go-callvis

生成调用关系图

```shell
go-callvis -debug -group "pkg,type" -focus "path/go/analyze" ./ | dot -T png -o output.png
```

不包含标准包

```shell
go-callvis -debug -nostd -group "pkg,type" -focus "path/go/analyze" ./ | dot -T png -o output.png
```

