### tcpdump

抓包保存

```shell
# 其中 -w 表示将抓的包保存到文件
# -C 表示当文件超过多大之后开始切分文件
# -Z 表示什么样的用户可以打开保存的cap文件
tcpdump -i any port 443 -w temp.cap -C 1 -Z root

```

后台抓包

```shell
nohup tcpdump -i any port 443 -w temp.cap -C 1 -Z root &
# 结束抓包最好发 SIGINT 信号
kill -2 pid
```

查看cap文件

```shell
tcpdump -r temp.cap -XX -nn
```

抓发送数据的包

```shell
# tcp flags (man tcpdump)
tcpdump -i eth0 -nn -xx tcp and dst host www.baidu.com and dst port 443 and 'tcp[tcpflags] & tcp-push != 0'
# output
15:43:50.940523 IP xx.xx.xx.xx.58566 > xx.xx.xx.xx.1024: Flags [P.], seq 36:42, ack 37, win 229, options [nop,nop,TS val 1400674059 ecr 1649552270], length 6
        0x0000:  fa16 3efc 32df fa16 3eaa 8f3d 0800 4500
        0x0010:  003a 9997 4000 3506 ed36 7274 40b3 0a16
        0x0020:  01b3 e4c6 0400 a23c 179a aeba 198b 8018
        0x0030:  00e5 c67e 0000 0101 080a 537c 970b 6252
        0x0040:  2b8e 0000 0006 0d0a
```

其中length代表tcp中的数据长度



