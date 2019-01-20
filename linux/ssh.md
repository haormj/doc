## ssh

配置ssh免密登录

```shell
# 生成本机公私钥
ssh-keygen -t rsa 
# 将本机公钥拷贝到要登录对方用户的主目录下面的 .ssh/authorized_keys
# 目标主机要求
chmod 700 .ssh -R
chmod 600 .ssh/authorized_keys
# 注意有时候配置完成之后还不能够使用使用,是因为客户端的配置,客户端默认使用的私钥中
# 没有包含你生成的私钥
ssh -i /path/to/gen_private root@host
```

多私钥管理

```shell
# 在用户目录的.ssh目录中创建config文件,以root为例 
cat > /root/.ssh/config << EOF
Host gitlab.com
    HostName gitlab.com
    IdentityFile /path/to/idfile

Host github.com
    HostName github.com
    IdentityFile /path/to/idfile
EOF
```
ssh_exchange_identification: read: Connection reset by peer

```shell
vi /etc/hosts.allow
# add
sshd: ALL
# restart
systemctl restart sshd
```

