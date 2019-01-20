### mysql

安装mysql客户端在centos7上


https://dev.mysql.com/downloads/repo/yum/

下载完成之后,通过rpm安装

```shell
wget https://dev.mysql.com/get/mysql80-community-release-el7-1.noarch.rpm
rpm -Uvh mysql80-community-release-el7-1.noarch.rpm
yum install mysql-community-client.x86_64 
```

mysql binlog文档

https://blog.csdn.net/leshami/article/details/39801867



mysql client登录数据库

```
mysql -h 127.0.0.1 -P 3306 -u xx -p xx
```

mysqlbinlog文档

https://blog.csdn.net/leshami/article/details/41962243

gtid

http://www.ywnds.com/?p=3898

centos安装mysql

https://dev.mysql.com/doc/mysql-yum-repo-quick-guide/en/

golang消费binlog

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
)

func main() {

	canalDataFile := "./canal.data"

	opt := canal.NewDefaultConfig()
	opt.Addr = "mysql.haormj.com:3306"
	opt.ReadTimeout = time.Minute
	opt.User = "mysql_slave"
	opt.Password = "Haormj@2018"
	c, err := canal.NewCanal(opt)
	if err != nil {
		log.Fatalln(err)
	}
	c.SetEventHandler(new(TestHandler))

	// obtain GTID
	// 1. from file
	// 2. from master
	var gtidSet mysql.GTIDSet
	canalDataBytes, err := ioutil.ReadFile(canalDataFile)
	if err != nil || len(canalDataBytes) == 0 {
		log.Println("Obtain GTIDSet from master")
		gtidSet, err = c.GetMasterGTIDSet()
	} else {
		log.Println("Obtain GTIDSet from local")
		gtidSet, err = mysql.ParseMysqlGTIDSet(string(canalDataBytes))
	}
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("start from GTID %v\n", gtidSet)
	go func() {
		if err := c.StartFromGTID(gtidSet); err != nil {
			log.Println(err)
			return
		}
	}()

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Println(err)
			return
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	select {
	case sig := <-ch:
		log.Printf("receive signal %v\n", sig)
		c.Close()
		gtidSet := c.SyncedGTIDSet()
		if err := ioutil.WriteFile(canalDataFile, []byte(gtidSet.String()), 0644); err != nil {
			log.Println("writer file error")
			return
		}
	}
}

type TestHandler struct {
	*canal.DummyEventHandler
}

func (t *TestHandler) OnRow(e *canal.RowsEvent) error {
	log.Println(e)
	return nil
}
```













