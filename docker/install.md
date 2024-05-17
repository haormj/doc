## 安装docker

阿里巴巴开源镜像站点

https://opsx.alibaba.com/

阿里云docker安装帮助文档

https://yq.aliyun.com/articles/110806

### 1.7.1

```
Hi everyone,

We have just pushed 1.7.1 to get.docker.com. You can download it with the following:

Ubuntu/Debian: curl -sSL https://get.docker.com | sh

Linux 64bit binary: https://get.docker.com/builds/Linux/x86_64/docker-1.7.1

Darwin/OSX 64bit client binary: https://get.docker.com/builds/Darwin/x86_64/docker-1.7.1

Darwin/OSX 32bit client binary: https://get.docker.com/builds/Darwin/i386/docker-1.7.1

Linux 64bit tgz: https://get.docker.com/builds/Linux/x86_64/docker-1.7.1.tgz

Windows 64bit client binary: https://get.docker.com/builds/Windows/x86_64/docker-1.7.1.exe

Windows 32bit client binary: https://get.docker.com/builds/Windows/i386/docker-1.7.1.exe

Centos 6/RHEL 6: https://get.docker.com/rpm/1.7.1/centos-6/RPMS/x86_64/docker-engine-1.7.1-1.el6.x86_64.rpm

Centos 7/RHEL 7: https://get.docker.com/rpm/1.7.1/centos-7/RPMS/x86_64/docker-engine-1.7.1-1.el7.centos.x86_64.rpm

Fedora 20: https://get.docker.com/rpm/1.7.1/fedora-20/RPMS/x86_64/docker-engine-1.7.1-1.fc20.x86_64.rpm

Fedora 21: https://get.docker.com/rpm/1.7.1/fedora-21/RPMS/x86_64/docker-engine-1.7.1-1.fc21.x86_64.rpm

Fedora 22: https://get.docker.com/rpm/1.7.1/fedora-22/RPMS/x86_64/docker-engine-1.7.1-1.fc22.x86_64.rpm

Check the changelog in https://github.com/docker/docker/blob/a35c4988e5f67cc4c8454f2e1afdda59e04f9542/CHANGELOG.md

Thanks to everyone that helped make this release possible!

Cheers,
```

https://groups.google.com/g/docker-dev/c/zfxGDRPEN7A

### 二进制安装

centos6 

如果是提示cgroup-fs的问题
```
service cgconfig start
```

如果没有自动创建 docker 网桥可以如下
```
ip link add name docker0 type bridge
ip addr add dev docker0 172.17.0.1/16
```


