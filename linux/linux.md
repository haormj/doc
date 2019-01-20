## linux

修改linux时区

```shell
 cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
```

关闭linux交换分区

```shell
# 查看当前的交换分区
cat /proc/swaps
# 关闭所有交换设备和文件
swapoff -a
```

测试udp端口是否可用

```shell
nc -vuz xx.xx.xx.xx 4500
# Connection to xx.xx.xx.xx 4500 port [udp/ipsec-nat-t] succeeded!
```

解压.gz文件

```shell
gzip -d filename
```

复制保留目录结构

```shell
cp --parents src dst
```

查看二进制依赖的动态链接库

```shell
ldd /bin/bash
```

查看dns信息

```shell
cat /etc/resolv.conf 
```

查看默认网关信息

```shell
ip route show
```

查看系统最大文件描述符

```shell
ulimit -n
```

统计tcp连接状态

```shell
while true; do netstat -apnt | awk '{print $6}' | sort | uniq -c; sleep 1; done
```

