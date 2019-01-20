## iptables

### 简介

网络报通过iptables的流程:

```
                               XXXXXXXXXXXXXXXXXX
                             XXX     Network    XXX
                               XXXXXXXXXXXXXXXXXX
                                       +
                                       |
                                       v
 +-------------+              +------------------+
 |table: filter| <---+        | table: nat       |
 |chain: INPUT |     |        | chain: PREROUTING|
 +-----+-------+     |        +--------+---------+
       |             |                 |
       v             |                 v
 [local process]     |           ****************          +--------------+
       |             +---------+ Routing decision +------> |table: filter |
       v                         ****************          |chain: FORWARD|
****************                                           +------+-------+
Routing decision                                                  |
****************                                                  |
       |                                                          |
       v                        ****************                  |
+-------------+       +------>  Routing decision  <---------------+
|table: nat   |       |         ****************
|chain: OUTPUT|       |               +
+-----+-------+       |               |
      |               |               v
      v               |      +-------------------+
+--------------+      |      | table: nat        |
|table: filter | +----+      | chain: POSTROUTING|
|chain: OUTPUT |             +--------+----------+
+--------------+                      |
                                      v
                               XXXXXXXXXXXXXXXXXX
                             XXX    Network     XXX
                               XXXXXXXXXXXXXXXXXX
```

#### 常用格式

```shell
iptables -t 表名 <-A/I/D/R> 规则链名 [规则号] <-i/o 网卡名> -p 协议名 <-s 源IP/源子网> --sport 源端口 <-d 目标IP/目标子网> --dport 目标端口 -j 动作
```

#### 插入

```shell
iptables -I INPUT -p tcp -s 192.168.1.2 -m tcp --dport 80 -j DROP 
```

#### 限制访问速率

```
iptables -A INPUT -i eth0 -p tcp -m limit --limit 20/s --limit-burst 5 -j ACCEPT
iptables -A INPUT -i eth0 -p tcp -j DROP
```

具体关于limit的内容可以通过`man iptables-extensions`

`--limit`表示的意思为每秒接受20个包,超过这个数量之后,则不由这条规则处理.其实标识每秒产生20个令牌.每当网络报匹配一次该规则,则消耗一个令牌.

`--limit-burst`表示初始化可以接受多少个包,比如刚开始的时候,你可以发现其能每秒能够接受25个包,之后则每秒能够接受20个包.更加简单的理解就是初始化的令牌池,当没有访问的时候,这个池子最大为5,之后当来访问的时候,则可以直接消耗池中令牌的数量,当令牌的数量消耗完的时候,就只能依靠`--limit`中产生令牌的数量.每个`limit`的周期内,`limit-burst`的值都尝试发生变化,只有当本次`limit`产生的令牌未被消耗完的时候,则直接增加到`burst`上,但不能超过其最大值.

需要注意采用`limit`限制,一定要配合类似其下面的规则使用,当不满足第一条规则的时候,则由后续的规则处理.从而需要后续的规则将其丢弃掉.

#### target

DROP: 直接丢弃这个包,终止匹配当前链,开始匹配下个链

REJECT: 丢弃并给对方返回错误包,终止匹配当前链,开始匹配下个链

### 参考

[Iptables](https://wiki.archlinux.org/index.php/Iptables_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87))

[iptables](http://wangchujiang.com/linux-command/c/iptables.html)