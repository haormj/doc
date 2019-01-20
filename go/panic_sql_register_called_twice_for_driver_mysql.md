## panic: sql: Register called twice for driver mysql

### 问题复现

#### 代码

目录结构

```
./
├── hello
│   ├── hello.go
│   └── vendor
│       └── github.com
│           └── go-sql-driver
│               └── mysql
│                   ├── appengine.go
│                   ├── AUTHORS
│                   ├── buffer.go
│                   ├── CHANGELOG.md
│                   ├── collations.go
│                   ├── connection.go
│                   ├── const.go
│                   ├── CONTRIBUTING.md
│                   ├── driver.go
│                   ├── dsn.go
│                   ├── errors.go
│                   ├── infile.go
│                   ├── ISSUE_TEMPLATE.md
│                   ├── LICENSE
│                   ├── packets.go
│                   ├── PULL_REQUEST_TEMPLATE.md
│                   ├── README.md
│                   ├── result.go
│                   ├── rows.go
│                   ├── statement.go
│                   ├── transaction.go
│                   └── utils.go
└── main.go
```

hello.go

```go
package hello

import (
	_ "github.com/go-sql-driver/mysql"
)
```

main.go

```go
package main

import (
	"fmt"

	_ "haormj.com/test/hello"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("hello")
}
```

#### 现象分析

##### 流程

1. main包导入hello包
2. hello包执行导入go-sql-driver包,使用的包为vendor的包
3. 执行verdor go-sql-driver包中的init方法
4. main执行导入go-sql-driver包,使用包为GOPATH中的包
5. 执行GOPATH go-sql-driver包中的init方法

##### go-sql-driver init源码

```go
func init() {
	sql.Register("mysql", &MySQLDriver{})
}
```

其中sql包为golang官方包`"database/sql"`

##### 结论

`sql.Register("mysql", &MySQLDriver{})`被执行两次导致panic

### 解决方案

`github.com/go-sql-driver/mysql`这个包最好不要加到vendor中.这样就不会导致`init`方法被执行两次.

### 参考

1. [`main`函数和`init`函数](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.3.md)