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

查看当前动态链接库信息

```shell
ldconfig -p
```

手动配置动态链接库
```shell
# 在/etc/ld.so.conf.d目录中创建对应文件,然后运行
# 重建缓存
ldconfig
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

计算字符串长度

```shell
echo "e2241ca1xf7caeea21111e1151211369f67f5b" | wc -L
```

快速打开文件管理器

```shell
nautilus /path
```

创建软链接

```shell
ln -s target link_name
```

搜狗拼音不能使用

```shell
# https://www.findhao.net/res/786
fcitx -r
```

ubuntu上制作window启动盘
https://my.oschina.net/u/2306127/blog/1796090
```

```

