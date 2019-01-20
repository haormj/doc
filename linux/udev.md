## udev

修改网卡名称

```
# udev 配置文件所在目录
# /etc/udev/rules.d
# /run/udev/rules.d
# /usr/lib/udev/rules.d
# 其中优先级从高到低,这里直接在/etc下设置
# 前提在linux内核2.6之后
# 检查/etc/udev/rules.d目录下是否存在70-persistent-net.rules,若存在则修改,否则创建
touch /etc/udev/rules.d/70-snap.core.rules
```

文件内容

```
SUBSYSTEM=="net", ACTION=="add", DRIVERS=="?*", ATTR{address}=="06:6a:6f:6a:26:1a", NAME="eth0"
```

其中NAME就是你想要修改的网卡名称,ATTR{address}写mac地址



关闭有规律的分配网卡名称,这个只是记录

```
/etc/default/grub 
GRUB_CMDLINE_LINUX 加上 net.ifnames=0
grub2-mkconfig -o /boot/grub2/grub.cfg
```

也可以

```
ln -s /dev/null 80-net-name-slot.rules
```













